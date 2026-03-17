// DYNAMICS-8 Personality Framework - Go Reference Implementation
// Copyright (c) 2026 Kronaxis Limited. All rights reserved.
// Licensed under BSL 1.1. See LICENSE file.
//
// Created by Jason Duke, Kronaxis Limited
// https://kronaxis.co.uk/dynamics
// Patent: UK Patent Application C (GB 2605150.8)

// Package dynamics provides the canonical Go reference implementation of the
// DYNAMICS-8 personality framework: eight orthogonal dimensions that describe
// an individual's behavioural tendencies on a continuous 0.0-1.0 scale.
package dynamics

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// DimensionInfo describes a single DYNAMICS-8 dimension and its four facets.
type DimensionInfo struct {
	Name        string
	Description string
	Facets      [4]string
}

// Dimensions maps each letter code to its definition.
var Dimensions = map[string]DimensionInfo{
	"D": {Name: "Discipline", Description: "Self-regulation, planning, routine adherence",
		Facets: [4]string{"Organisation", "Diligence", "Perfectionism", "Prudence"}},
	"Y": {Name: "Yielding", Description: "Agreeableness, compliance, conflict avoidance",
		Facets: [4]string{"Patience", "Tolerance", "Flexibility", "Persuadability"}},
	"N": {Name: "Novelty", Description: "Openness to new experiences, curiosity",
		Facets: [4]string{"Aesthetic sense", "Inquisitiveness", "Creativity", "Unconventionality"}},
	"A": {Name: "Acuity", Description: "Digital fluency, platform nativeness, research tendency",
		Facets: [4]string{"Platform nativeness", "Content creation", "Privacy awareness", "Tech adoption"}},
	"M": {Name: "Mercuriality", Description: "Emotional volatility, anxiety proneness",
		Facets: [4]string{"Anxiety", "Emotional reactivity", "Sentimentality", "Dependence"}},
	"I": {Name: "Impulsivity", Description: "Decision speed, spontaneous action",
		Facets: [4]string{"Delay discounting", "Sensation seeking", "Snap decision", "Boredom susceptibility"}},
	"C": {Name: "Candour", Description: "Honesty, ethical concern, transparency",
		Facets: [4]string{"Sincerity", "Fairness", "Modesty", "Materialism inverse"}},
	"S": {Name: "Sociability", Description: "Social engagement, group orientation",
		Facets: [4]string{"Social boldness", "Liveliness", "Social esteem", "Gregariousness"}},
}

// DimensionOrder is the canonical ordering of dimension letters.
var DimensionOrder = [8]string{"D", "Y", "N", "A", "M", "I", "C", "S"}

// DynamicsProfile holds the 8 DYNAMICS dimension scores (0.0-1.0 each).
type DynamicsProfile struct {
	D      float64               `json:"D"` // Discipline
	Y      float64               `json:"Y"` // Yielding
	N      float64               `json:"N"` // Novelty
	A      float64               `json:"A"` // Acuity
	M      float64               `json:"M"` // Mercuriality
	I      float64               `json:"I"` // Impulsivity
	C      float64               `json:"C"` // Candour
	S      float64               `json:"S"` // Sociability
	Facets map[string][4]float64 `json:"facets,omitempty"`
}

// scores returns all eight dimension values in canonical (D-Y-N-A-M-I-C-S) order.
func (p *DynamicsProfile) scores() [8]float64 {
	return [8]float64{p.D, p.Y, p.N, p.A, p.M, p.I, p.C, p.S}
}

// Validate checks that every dimension score falls within the [0.0, 1.0] range.
// Also validates facet scores if present.
func (p *DynamicsProfile) Validate() error {
	vals := p.scores()
	for i, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) || v < 0.0 || v > 1.0 {
			return fmt.Errorf("dimension %s (%s) out of range: %.4f (must be 0.0-1.0)",
				DimensionOrder[i], Dimensions[DimensionOrder[i]].Name, v)
		}
	}
	for dim, facets := range p.Facets {
		if _, ok := Dimensions[dim]; !ok {
			return fmt.Errorf("unknown facet dimension: %s", dim)
		}
		for i, f := range facets {
			if f < 0.0 || f > 1.0 {
				return fmt.Errorf("facet %s[%d] out of range: %.4f", dim, i, f)
			}
		}
	}
	return nil
}

// DimensionLabel returns a human-readable band label for a score in [0.0, 1.0].
//
//	[0.00, 0.20) -> "very low"
//	[0.20, 0.40) -> "low"
//	[0.40, 0.60) -> "moderate"
//	[0.60, 0.80) -> "high"
//	[0.80, 1.00] -> "very high"
func DimensionLabel(score float64) string {
	switch {
	case score >= 0.80:
		return "very high"
	case score >= 0.60:
		return "high"
	case score >= 0.40:
		return "moderate"
	case score >= 0.20:
		return "low"
	default:
		return "very low"
	}
}

// Summary returns a human-readable text representation of the profile.
// Example: "D: Discipline = 0.71 (high), Y: Yielding = 0.55 (moderate), ..."
func (p *DynamicsProfile) Summary() string {
	vals := p.scores()
	parts := make([]string, 8)
	for i, v := range vals {
		dim := DimensionOrder[i]
		parts[i] = fmt.Sprintf("%s: %s = %.2f (%s)",
			dim, Dimensions[dim].Name, v, DimensionLabel(v))
	}
	return strings.Join(parts, ", ")
}

// Octant returns an 8-segment code representing the high/low classification of
// each dimension. Each segment is the dimension letter followed by H (>= 0.5)
// or L (< 0.5), separated by hyphens. Example: "DH-YL-NH-AL-MH-IL-CH-SL".
func (p *DynamicsProfile) Octant() string {
	vals := p.scores()
	parts := make([]string, 8)
	for i, v := range vals {
		suffix := "L"
		if v >= 0.5 {
			suffix = "H"
		}
		parts[i] = DimensionOrder[i] + suffix
	}
	return strings.Join(parts, "-")
}

// ToJSON serialises the profile to compact JSON.
func (p *DynamicsProfile) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// FromJSON deserialises a JSON object into the profile and validates the result.
func (p *DynamicsProfile) FromJSON(data []byte) error {
	if err := json.Unmarshal(data, p); err != nil {
		return fmt.Errorf("dynamics: invalid JSON: %w", err)
	}
	if err := p.Validate(); err != nil {
		return fmt.Errorf("dynamics: validation failed after deserialisation: %w", err)
	}
	return nil
}

// ---- Beta distribution sampling ----

// gammaVariate produces a sample from the Gamma(alpha, 1) distribution using
// Marsaglia and Tsang's method for alpha >= 1 and a boosting transform for
// alpha < 1. This avoids any dependency on external statistical libraries.
func gammaVariate(rng *rand.Rand, alpha float64) float64 {
	if alpha < 1.0 {
		// Boost: Gamma(alpha,1) = Gamma(alpha+1,1) * U^(1/alpha)
		return gammaVariate(rng, alpha+1.0) * math.Pow(rng.Float64(), 1.0/alpha)
	}
	// Marsaglia and Tsang's method (alpha >= 1)
	d := alpha - 1.0/3.0
	c := 1.0 / math.Sqrt(9.0*d)
	for {
		var x, v float64
		for {
			x = rng.NormFloat64()
			v = 1.0 + c*x
			if v > 0 {
				break
			}
		}
		v = v * v * v
		u := rng.Float64()
		if u < 1.0-0.0331*(x*x)*(x*x) {
			return d * v
		}
		if math.Log(u) < 0.5*x*x+d*(1.0-v+math.Log(v)) {
			return d * v
		}
	}
}

// betaVariate produces a sample from the Beta(a, b) distribution.
// beta(a,b) = g1 / (g1 + g2) where g1 ~ Gamma(a,1) and g2 ~ Gamma(b,1).
func betaVariate(rng *rand.Rand, a, b float64) float64 {
	g1 := gammaVariate(rng, a)
	g2 := gammaVariate(rng, b)
	return g1 / (g1 + g2)
}

// ---- Profile generation ----

// GenerateProfile produces a random DYNAMICS-8 profile where each dimension is
// drawn independently from a Beta(2,2) distribution. This yields a bell-shaped
// spread centred on 0.5 with moderate variance, avoiding the clustering at
// extremes that a uniform distribution would produce.
func GenerateProfile() DynamicsProfile {
	rng := rand.New(rand.NewSource(rand.Int63()))
	return DynamicsProfile{
		D: betaVariate(rng, 2, 2),
		Y: betaVariate(rng, 2, 2),
		N: betaVariate(rng, 2, 2),
		A: betaVariate(rng, 2, 2),
		M: betaVariate(rng, 2, 2),
		I: betaVariate(rng, 2, 2),
		C: betaVariate(rng, 2, 2),
		S: betaVariate(rng, 2, 2),
	}
}

// ---- Compatibility ----

// compatWeights defines the per-dimension weights for compatibility scoring.
// Positive weights reward similarity (small difference between profiles);
// negative weights reward complementarity (large difference).
var compatWeights = [8]float64{
	0.8,  // D: Discipline (similarity)
	-0.3, // Y: Yielding (complementarity)
	0.6,  // N: Novelty (similarity)
	0.5,  // A: Acuity (similarity)
	-0.4, // M: Mercuriality (complementarity)
	0.2,  // I: Impulsivity (similarity)
	0.9,  // C: Candour (similarity)
	0.4,  // S: Sociability (similarity)
}

// CompatibilityScore computes a weighted similarity/complementarity score
// between two profiles. The result is normalised to [0.0, 1.0] where 1.0
// indicates maximum compatibility.
func CompatibilityScore(a, b DynamicsProfile) float64 {
	aVals := a.scores()
	bVals := b.scores()

	var totalWeight float64
	var score float64
	for i := 0; i < 8; i++ {
		diff := math.Abs(aVals[i] - bVals[i])
		w := math.Abs(compatWeights[i])
		totalWeight += w
		if compatWeights[i] >= 0 {
			// Positive weight: similarity is good (small diff -> high contribution)
			score += w * (1.0 - diff)
		} else {
			// Negative weight: complementarity is good (large diff -> high contribution)
			score += w * diff
		}
	}
	if totalWeight == 0 {
		return 0.5
	}
	return score / totalWeight
}

// ---- Derived behavioural indicators ----

// DeriveIncomeBand estimates a relative income-band indicator from the profile.
// Higher Discipline and Novelty correlate with higher earning potential.
// Returns a value in [0.0, 1.0].
func DeriveIncomeBand(p DynamicsProfile) float64 {
	return clamp01(p.D*0.6 + p.N*0.4)
}

// DeriveSpendingPattern estimates spending prudence from the profile.
// Higher Candour, Discipline, and lower Impulsivity indicate more deliberate spending.
// Returns a value in [0.0, 1.0] where 1.0 is highly prudent.
func DeriveSpendingPattern(p DynamicsProfile) float64 {
	return clamp01(p.C*0.3 + p.D*0.4 + (1-p.I)*0.3)
}

// DeriveRiskTolerance estimates willingness to take risks.
// Higher Yielding, Impulsivity, and lower Mercuriality indicate greater risk tolerance.
// Returns a value in [0.0, 1.0].
func DeriveRiskTolerance(p DynamicsProfile) float64 {
	return clamp01(p.Y*0.25 + p.I*0.40 + (1-p.M)*0.35)
}

// DerivePoliticalLean estimates a left-right political tendency from the profile.
// Returns a value in [-1.0, +1.0] where -1.0 is strongly left-leaning and
// +1.0 is strongly right-leaning. The mapping is intentionally coarse: real
// political orientation is far more complex than any eight-dimensional model
// can capture, but this provides a useful heuristic for persona generation.
//
// Heuristics applied:
//   - High Novelty and low Discipline push left (openness, less tradition)
//   - High Discipline and low Yielding push right (order, firmness)
//   - High Candour moderates towards centre (fairness, modesty)
//   - High Sociability has a mild leftward pull (communal orientation)
func DerivePoliticalLean(p DynamicsProfile) float64 {
	// Start at centre
	lean := 0.0

	// Novelty pushes left; its absence pushes right
	lean -= (p.N - 0.5) * 0.8

	// Discipline pushes right; its absence pushes left
	lean += (p.D - 0.5) * 0.6

	// Low Yielding (stubbornness, firmness) pushes right
	lean -= (p.Y - 0.5) * 0.4

	// Candour (fairness, modesty) pulls towards centre
	lean *= 1.0 - (p.C * 0.2)

	// Sociability has a mild leftward pull
	lean -= (p.S - 0.5) * 0.2

	// Clamp to [-1.0, +1.0]
	if lean < -1.0 {
		return -1.0
	}
	if lean > 1.0 {
		return 1.0
	}
	return lean
}

// clamp01 restricts a value to the [0.0, 1.0] range.
func clamp01(v float64) float64 {
	if v < 0.0 {
		return 0.0
	}
	if v > 1.0 {
		return 1.0
	}
	return v
}
