// DYNAMICS-8 Personality Framework - Go Test Suite
// Copyright (c) 2026 Kronaxis Limited. All rights reserved.
// Licensed under BSL 1.1. See LICENSE-BSL.

package dynamics

import (
	"encoding/json"
	"math"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// TestDimensions: verify the Dimensions constant is complete and well-formed
// ---------------------------------------------------------------------------

func TestEightDimensionsPresent(t *testing.T) {
	if len(Dimensions) != 8 {
		t.Fatalf("expected 8 dimensions, got %d", len(Dimensions))
	}
	for _, ch := range "DYNAMICS" {
		key := string(ch)
		if _, ok := Dimensions[key]; !ok {
			t.Errorf("dimension %q not found in Dimensions map", key)
		}
	}
}

func TestEachDimensionHasFourFacets(t *testing.T) {
	for key, info := range Dimensions {
		if len(info.Facets) != 4 {
			t.Errorf("dimension %s should have 4 facets, got %d", key, len(info.Facets))
		}
		for i, f := range info.Facets {
			if f == "" {
				t.Errorf("dimension %s facet[%d] is empty", key, i)
			}
		}
	}
}

func TestDimensionNames(t *testing.T) {
	expected := map[string]string{
		"D": "Discipline", "Y": "Yielding", "N": "Novelty", "A": "Acuity",
		"M": "Mercuriality", "I": "Impulsivity", "C": "Candour", "S": "Sociability",
	}
	for key, name := range expected {
		info, ok := Dimensions[key]
		if !ok {
			t.Fatalf("dimension %s missing", key)
		}
		if info.Name != name {
			t.Errorf("dimension %s: expected name %q, got %q", key, name, info.Name)
		}
	}
}

// ---------------------------------------------------------------------------
// TestDynamicsProfile: creation, validation, serialisation
// ---------------------------------------------------------------------------

func TestDefaultProfileIsValid(t *testing.T) {
	p := DynamicsProfile{D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	if err := p.Validate(); err != nil {
		t.Errorf("default profile should validate, got: %v", err)
	}
}

func TestOutOfRangeFailsValidation(t *testing.T) {
	p := DynamicsProfile{D: 1.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	if err := p.Validate(); err == nil {
		t.Error("profile with D=1.5 should fail validation")
	}
}

func TestNegativeFailsValidation(t *testing.T) {
	p := DynamicsProfile{D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: -0.1, I: 0.5, C: 0.5, S: 0.5}
	if err := p.Validate(); err == nil {
		t.Error("profile with M=-0.1 should fail validation")
	}
}

func TestJSONRoundTrip(t *testing.T) {
	original := DynamicsProfile{
		D: 0.71, Y: 0.55, N: 0.83, A: 0.90, M: 0.42, I: 0.35, C: 0.65, S: 0.78,
	}
	data, err := original.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	var restored DynamicsProfile
	if err := restored.FromJSON(data); err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	origScores := original.scores()
	restScores := restored.scores()
	for i := 0; i < 8; i++ {
		if math.Abs(origScores[i]-restScores[i]) > 1e-6 {
			t.Errorf("dimension %s: original %.6f != restored %.6f",
				DimensionOrder[i], origScores[i], restScores[i])
		}
	}
}

func TestJSONKeysAreSingleLetters(t *testing.T) {
	p := DynamicsProfile{D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	data, err := p.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("could not unmarshal JSON: %v", err)
	}

	for _, ch := range "DYNAMICS" {
		key := string(ch)
		if _, ok := raw[key]; !ok {
			t.Errorf("JSON key %q not found in serialised output", key)
		}
	}
}

func TestSummaryContainsDimensionNames(t *testing.T) {
	p := DynamicsProfile{D: 0.9, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.1}
	summary := p.Summary()
	if !strings.Contains(summary, "Discipline") {
		t.Error("summary should contain 'Discipline'")
	}
	if !strings.Contains(summary, "Sociability") {
		t.Error("summary should contain 'Sociability'")
	}
}

func TestDimensionLabelRanges(t *testing.T) {
	tests := []struct {
		score    float64
		expected string
	}{
		{0.05, "very low"},
		{0.10, "very low"},
		{0.19, "very low"},
		{0.20, "low"},
		{0.30, "low"},
		{0.39, "low"},
		{0.40, "moderate"},
		{0.50, "moderate"},
		{0.59, "moderate"},
		{0.60, "high"},
		{0.70, "high"},
		{0.79, "high"},
		{0.80, "very high"},
		{0.90, "very high"},
		{0.95, "very high"},
		{1.00, "very high"},
	}
	for _, tt := range tests {
		got := DimensionLabel(tt.score)
		if got != tt.expected {
			t.Errorf("DimensionLabel(%.2f) = %q, want %q", tt.score, got, tt.expected)
		}
	}
}

func TestOctantFormat(t *testing.T) {
	p := DynamicsProfile{D: 0.8, Y: 0.3, N: 0.7, A: 0.2, M: 0.6, I: 0.4, C: 0.9, S: 0.1}
	octant := p.Octant()
	expected := "DH-YL-NH-AL-MH-IL-CH-SL"
	if octant != expected {
		t.Errorf("Octant() = %q, want %q", octant, expected)
	}
}

func TestOctantBoundary(t *testing.T) {
	// Exactly 0.5 should be classified as High
	p := DynamicsProfile{D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	octant := p.Octant()
	expected := "DH-YH-NH-AH-MH-IH-CH-SH"
	if octant != expected {
		t.Errorf("Octant() at 0.5 boundary = %q, want %q", octant, expected)
	}
}

// ---------------------------------------------------------------------------
// TestGeneration: verify profile generation
// ---------------------------------------------------------------------------

func TestGeneratedProfileIsValid(t *testing.T) {
	for i := 0; i < 100; i++ {
		p := GenerateProfile()
		if err := p.Validate(); err != nil {
			data, _ := p.ToJSON()
			t.Fatalf("generated profile %d failed validation: %v (profile: %s)", i, err, string(data))
		}
	}
}

func TestOctantDistribution(t *testing.T) {
	// Generate 200 profiles and check no single octant exceeds 30%.
	octants := make(map[string]int)
	const n = 200
	for i := 0; i < n; i++ {
		p := GenerateProfile()
		o := p.Octant()
		octants[o]++
	}
	maxCount := 0
	for _, count := range octants {
		if count > maxCount {
			maxCount = count
		}
	}
	maxPct := float64(maxCount) / float64(n)
	if maxPct >= 0.30 {
		t.Errorf("single octant should not exceed 30%% of generated profiles, got %.1f%% (%d/%d)",
			maxPct*100, maxCount, n)
	}
}

func TestGeneratedProfilesHaveVariance(t *testing.T) {
	// Ensure generated profiles are not all identical.
	var first DynamicsProfile
	allSame := true
	for i := 0; i < 20; i++ {
		p := GenerateProfile()
		if i == 0 {
			first = p
		} else if p.D != first.D || p.Y != first.Y || p.N != first.N {
			allSame = false
		}
	}
	if allSame {
		t.Error("20 generated profiles were all identical; generation should produce variance")
	}
}

// ---------------------------------------------------------------------------
// TestCompatibility: verify compatibility scoring
// ---------------------------------------------------------------------------

func TestSelfCompatibilityIsHigh(t *testing.T) {
	for i := 0; i < 20; i++ {
		p := GenerateProfile()
		score := CompatibilityScore(p, p)
		if score < 0.5 {
			t.Errorf("self-compatibility should be >= 0.5, got %.4f", score)
		}
	}
}

func TestCompatibilityRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		a := GenerateProfile()
		b := GenerateProfile()
		score := CompatibilityScore(a, b)
		if score < 0.0 || score > 1.0 {
			t.Errorf("compatibility score out of [0,1] range: %.4f", score)
		}
	}
}

func TestCompatibilityIsSymmetric(t *testing.T) {
	for i := 0; i < 20; i++ {
		a := GenerateProfile()
		b := GenerateProfile()
		ab := CompatibilityScore(a, b)
		ba := CompatibilityScore(b, a)
		if math.Abs(ab-ba) > 1e-10 {
			t.Errorf("compatibility should be symmetric: score(a,b)=%.6f != score(b,a)=%.6f", ab, ba)
		}
	}
}

func TestIdenticalProfilesMaximalCompatibility(t *testing.T) {
	// Two identical profiles should have maximum compatibility on similarity
	// dimensions and minimum on complementarity dimensions. Self-score should
	// be consistently high due to weight structure.
	p := DynamicsProfile{D: 0.7, Y: 0.5, N: 0.6, A: 0.8, M: 0.3, I: 0.5, C: 0.8, S: 0.6}
	score := CompatibilityScore(p, p)
	if score < 0.5 {
		t.Errorf("identical profile compatibility should be >= 0.5, got %.4f", score)
	}
}

// ---------------------------------------------------------------------------
// TestDerivations: verify derived behavioural indicators
// ---------------------------------------------------------------------------

func TestDeriveIncomeBandRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		p := GenerateProfile()
		v := DeriveIncomeBand(p)
		if v < 0.0 || v > 1.0 {
			t.Errorf("DeriveIncomeBand out of [0,1] range: %.4f", v)
		}
	}
}

func TestDeriveSpendingPatternRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		p := GenerateProfile()
		v := DeriveSpendingPattern(p)
		if v < 0.0 || v > 1.0 {
			t.Errorf("DeriveSpendingPattern out of [0,1] range: %.4f", v)
		}
	}
}

func TestDeriveRiskToleranceRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		p := GenerateProfile()
		v := DeriveRiskTolerance(p)
		if v < 0.0 || v > 1.0 {
			t.Errorf("DeriveRiskTolerance out of [0,1] range: %.4f", v)
		}
	}
}

func TestDerivePoliticalLeanRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		p := GenerateProfile()
		lean := DerivePoliticalLean(p)
		if lean < -1.0 || lean > 1.0 {
			t.Errorf("DerivePoliticalLean out of [-1,1] range: %.4f", lean)
		}
	}
}

func TestHighDHighNTendsUpperIncome(t *testing.T) {
	// High D + high N should trend toward higher income band values.
	// With D=0.9 and N=0.9, DeriveIncomeBand = clamp01(0.9*0.6 + 0.9*0.4) = 0.9.
	// This is deterministic given fixed D and N, so every profile yields the same result.
	p := DynamicsProfile{D: 0.9, Y: 0.5, N: 0.9, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	v := DeriveIncomeBand(p)
	if v < 0.7 {
		t.Errorf("high D + high N should produce income band >= 0.7, got %.4f", v)
	}
}

func TestLowDLowNTendsLowerIncome(t *testing.T) {
	p := DynamicsProfile{D: 0.1, Y: 0.5, N: 0.1, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5}
	v := DeriveIncomeBand(p)
	if v > 0.3 {
		t.Errorf("low D + low N should produce income band <= 0.3, got %.4f", v)
	}
}

func TestHighImpulsivityLowSpendingPrudence(t *testing.T) {
	// High impulsivity should reduce spending prudence.
	prudent := DynamicsProfile{D: 0.8, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.1, C: 0.8, S: 0.5}
	impulsive := DynamicsProfile{D: 0.2, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.9, C: 0.2, S: 0.5}
	vPrudent := DeriveSpendingPattern(prudent)
	vImpulsive := DeriveSpendingPattern(impulsive)
	if vImpulsive >= vPrudent {
		t.Errorf("impulsive profile should have lower spending prudence: impulsive=%.4f >= prudent=%.4f",
			vImpulsive, vPrudent)
	}
}

func TestPoliticalLeanExtremes(t *testing.T) {
	// High novelty, low discipline, high sociability should lean left.
	leftish := DynamicsProfile{D: 0.1, Y: 0.8, N: 0.9, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.9}
	lean := DerivePoliticalLean(leftish)
	if lean >= 0 {
		t.Errorf("high N, low D, high S profile should lean left (negative), got %.4f", lean)
	}

	// High discipline, low novelty, low yielding should lean right.
	rightish := DynamicsProfile{D: 0.9, Y: 0.1, N: 0.1, A: 0.5, M: 0.5, I: 0.5, C: 0.2, S: 0.1}
	lean = DerivePoliticalLean(rightish)
	if lean <= 0 {
		t.Errorf("high D, low N, low Y profile should lean right (positive), got %.4f", lean)
	}
}

// ---------------------------------------------------------------------------
// TestEdgeCases: boundary conditions and corner cases
// ---------------------------------------------------------------------------

func TestValidateAllZeros(t *testing.T) {
	p := DynamicsProfile{}
	if err := p.Validate(); err != nil {
		t.Errorf("all-zero profile should be valid (0.0 is in range), got: %v", err)
	}
}

func TestValidateAllOnes(t *testing.T) {
	p := DynamicsProfile{D: 1.0, Y: 1.0, N: 1.0, A: 1.0, M: 1.0, I: 1.0, C: 1.0, S: 1.0}
	if err := p.Validate(); err != nil {
		t.Errorf("all-ones profile should be valid (1.0 is in range), got: %v", err)
	}
}

func TestValidateNaNFails(t *testing.T) {
	p := DynamicsProfile{D: math.NaN()}
	if err := p.Validate(); err == nil {
		t.Error("NaN dimension should fail validation")
	}
}

func TestValidateInfFails(t *testing.T) {
	p := DynamicsProfile{D: math.Inf(1)}
	if err := p.Validate(); err == nil {
		t.Error("+Inf dimension should fail validation")
	}
}

func TestFromJSONInvalidJSON(t *testing.T) {
	var p DynamicsProfile
	err := p.FromJSON([]byte("{invalid"))
	if err == nil {
		t.Error("FromJSON should fail on invalid JSON")
	}
}

func TestFromJSONOutOfRange(t *testing.T) {
	data := []byte(`{"D":2.0,"Y":0.5,"N":0.5,"A":0.5,"M":0.5,"I":0.5,"C":0.5,"S":0.5}`)
	var p DynamicsProfile
	err := p.FromJSON(data)
	if err == nil {
		t.Error("FromJSON should fail when deserialised values are out of range")
	}
}

func TestDimensionOrderMatchesDYNAMICS(t *testing.T) {
	expected := "DYNAMICS"
	for i, ch := range expected {
		if DimensionOrder[i] != string(ch) {
			t.Errorf("DimensionOrder[%d] = %q, want %q", i, DimensionOrder[i], string(ch))
		}
	}
}

func TestFacetValidation(t *testing.T) {
	p := DynamicsProfile{
		D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5,
		Facets: map[string][4]float64{
			"D": {0.6, 0.7, 0.8, 1.5}, // facet out of range
		},
	}
	if err := p.Validate(); err == nil {
		t.Error("profile with out-of-range facet should fail validation")
	}
}

func TestFacetUnknownDimension(t *testing.T) {
	p := DynamicsProfile{
		D: 0.5, Y: 0.5, N: 0.5, A: 0.5, M: 0.5, I: 0.5, C: 0.5, S: 0.5,
		Facets: map[string][4]float64{
			"X": {0.5, 0.5, 0.5, 0.5},
		},
	}
	if err := p.Validate(); err == nil {
		t.Error("profile with unknown facet dimension 'X' should fail validation")
	}
}

func TestClamp01(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-0.5, 0.0},
		{0.0, 0.0},
		{0.5, 0.5},
		{1.0, 1.0},
		{1.5, 1.0},
	}
	for _, tt := range tests {
		got := clamp01(tt.input)
		if got != tt.expected {
			t.Errorf("clamp01(%.1f) = %.1f, want %.1f", tt.input, got, tt.expected)
		}
	}
}

func TestDeriveIncomeBandExtremes(t *testing.T) {
	// D=0, N=0 should give 0.0
	low := DynamicsProfile{}
	if v := DeriveIncomeBand(low); v != 0.0 {
		t.Errorf("DeriveIncomeBand(all-zero) = %.4f, want 0.0", v)
	}
	// D=1, N=1 should give 1.0
	high := DynamicsProfile{D: 1.0, N: 1.0}
	if v := DeriveIncomeBand(high); v != 1.0 {
		t.Errorf("DeriveIncomeBand(D=1,N=1) = %.4f, want 1.0", v)
	}
}

func TestDeriveRiskToleranceExtremes(t *testing.T) {
	// Y=0, I=0, M=1 should give 0.0 (clamp)
	low := DynamicsProfile{M: 1.0}
	v := DeriveRiskTolerance(low)
	if v < 0.0 || v > 0.01 {
		t.Errorf("DeriveRiskTolerance with low risk factors should be near 0.0, got %.4f", v)
	}
}
