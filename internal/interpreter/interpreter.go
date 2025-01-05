package interpreter

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Posting => final interpretation: from -> to, how much, which asset
type Posting struct {
	Source      string
	Destination string
	Asset       string
	Amount      *big.Int
}

func EvaluateSendStmt(stmt *SendStmtAST) ([]Posting, error) {
	// parse amount => big.Int, or star => large
	isStar := (stmt.Amount == "*")
	var needed *big.Int
	if isStar {
		needed = big.NewInt(999999999) // naive large
	} else {
		tmp := new(big.Int)
		if _, ok := tmp.SetString(stmt.Amount, 10); !ok {
			return nil, fmt.Errorf("invalid amount: %s", stmt.Amount)
		}
		needed = tmp
	}

	srcParts, err := evalSource(stmt.Source, needed)
	if err != nil {
		return nil, err
	}
	totalSrc := sumPostingsAmt(srcParts)
	if !isStar && totalSrc.Cmp(needed) < 0 {
		return nil, fmt.Errorf("insufficient funds: need %s, have %s", needed, totalSrc)
	}
	var finalNeeded = needed
	if isStar {
		finalNeeded = totalSrc
	}

	destParts, err := evalDestination(stmt.Destination, finalNeeded)
	if err != nil {
		return nil, err
	}

	// Cross match
	var out []Posting
	si, di := 0, 0
	srcLeft := new(big.Int).Set(srcParts[si].Amount)
	dstLeft := new(big.Int).Set(destParts[di].Amount)
	for {
		if srcLeft.Sign() == 0 {
			si++
			if si >= len(srcParts) {
				break
			}
			srcLeft.Set(srcParts[si].Amount)
		}
		if dstLeft.Sign() == 0 {
			di++
			if di >= len(destParts) {
				break
			}
			dstLeft.Set(destParts[di].Amount)
		}
		if si >= len(srcParts) || di >= len(destParts) {
			break
		}
		m := minBigInt(srcLeft, dstLeft)
		out = append(out, Posting{
			Source:      srcParts[si].Source,
			Destination: destParts[di].Destination,
			Asset:       stmt.Asset,
			Amount:      new(big.Int).Set(m),
		})
		srcLeft.Sub(srcLeft, m)
		dstLeft.Sub(dstLeft, m)
	}
	return out, nil
}

func evalSource(s SourceAST, needed *big.Int) ([]Posting, error) {
	switch node := s.(type) {
	case SingleSrc:
		return []Posting{{
			Source:      node.Account,
			Destination: "",
			Asset:       "", // we fill later
			Amount:      new(big.Int).Set(needed),
		}}, nil
	case InorderSrc:
		left := new(big.Int).Set(needed)
		var result []Posting
		for _, sub := range node.Sources {
			if left.Sign() == 0 {
				break
			}
			chunk, err := evalSource(sub, left)
			if err != nil {
				return nil, err
			}
			sum := sumPostingsAmt(chunk)
			result = append(result, chunk...)
			left.Sub(left, sum)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unknown source AST")
	}
}

func evalDestination(d DestAST, total *big.Int) ([]Posting, error) {
	switch n := d.(type) {
	case SingleDest:
		return []Posting{{
			Source:      "",
			Destination: n.Account,
			Asset:       "",
			Amount:      new(big.Int).Set(total),
		}}, nil
	case AllotmentDest:
		return allotmentDestToPostings(n, total)
	default:
		return nil, fmt.Errorf("unknown destination AST")
	}
}

func allotmentDestToPostings(ad AllotmentDest, total *big.Int) ([]Posting, error) {
	// parse portion or remaining
	floors := make([]*big.Int, len(ad.Items))
	var sumFloors = big.NewInt(0)
	var portionSum big.Rat
	remIndex := -1
	for i, item := range ad.Items {
		if item.IsRemaining {
			remIndex = i
		} else {
			r, err := parsePortion(item.Portion)
			if err != nil {
				return nil, err
			}
			portionSum.Add(&portionSum, r)
		}
	}
	// first pass -> floor
	for i, item := range ad.Items {
		if item.IsRemaining {
			floors[i] = big.NewInt(0)
			continue
		}
		r, err := parsePortion(item.Portion)
		if err != nil {
			return nil, err
		}
		p := portionOf(total, r)
		floors[i] = p
		sumFloors.Add(sumFloors, p)
	}
	leftover := new(big.Int).Sub(total, sumFloors)
	// distribute leftover to portioned
	for i, item := range ad.Items {
		if leftover.Sign() == 0 {
			break
		}
		if item.IsRemaining {
			continue
		}
		floors[i].Add(floors[i], big.NewInt(1))
		leftover.Sub(leftover, big.NewInt(1))
	}
	// leftover to remaining
	if remIndex == -1 && leftover.Sign() != 0 {
		return nil, fmt.Errorf("No 'remaining' but leftover > 0")
	}
	if remIndex != -1 {
		floors[remIndex].Add(floors[remIndex], leftover)
	}
	var out []Posting
	for i, item := range ad.Items {
		if floors[i].Sign() == 0 {
			continue
		}
		out = append(out, Posting{
			Source:      "",
			Destination: item.Account,
			Asset:       "",
			Amount:      floors[i],
		})
	}
	return out, nil
}

func parsePortion(p string) (*big.Rat, error) {
	p = strings.TrimSpace(p)
	if strings.HasSuffix(p, "%") {
		ps := strings.TrimSuffix(p, "%")
		numf, err := strconv.ParseFloat(ps, 64)
		if err != nil {
			return nil, err
		}
		r := new(big.Rat).SetFloat64(numf)
		r.Mul(r, big.NewRat(1, 100))
		return r, nil
	}
	if strings.Contains(p, "/") {
		split := strings.Split(p, "/")
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid fraction %s", p)
		}
		num, err1 := strconv.ParseInt(strings.TrimSpace(split[0]), 10, 64)
		den, err2 := strconv.ParseInt(strings.TrimSpace(split[1]), 10, 64)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("cannot parse ratio: %s", p)
		}
		return new(big.Rat).SetFrac(big.NewInt(num), big.NewInt(den)), nil
	}
	// fallback => parse as rational
	r := new(big.Rat)
	if _, ok := r.SetString(p); !ok {
		return nil, fmt.Errorf("invalid portion: %s", p)
	}
	return r, nil
}

func portionOf(total *big.Int, r *big.Rat) *big.Int {
	tr := new(big.Rat).SetInt(total)
	prod := new(big.Rat).Mul(tr, r)
	fs := prod.FloatString(0) // floor
	bi := new(big.Int)
	bi.SetString(fs, 10)
	return bi
}

func sumPostingsAmt(pp []Posting) *big.Int {
	out := big.NewInt(0)
	for _, p := range pp {
		out.Add(out, p.Amount)
	}
	return out
}

func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) <= 0 {
		return a
	}
	return b
}
