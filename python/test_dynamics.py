# DYNAMICS-8 Personality Framework - Test Suite
# Copyright (c) 2026 Kronaxis Limited. All rights reserved.
# Licensed under BSL 1.1. See LICENSE-BSL.

"""Basic tests for the DYNAMICS-8 Python reference implementation."""

import json
import unittest
from dynamics import (
    DIMENSIONS,
    DynamicsProfile,
    generate_profile,
    compatibility_score,
    derive_income_band,
    derive_spending_pattern,
    derive_risk_tolerance,
    derive_political_lean,
)


class TestDimensions(unittest.TestCase):
    """Verify the DIMENSIONS constant is complete and well-formed."""

    def test_eight_dimensions_present(self):
        self.assertEqual(len(DIMENSIONS), 8)
        for key in "DYNAMICS":
            self.assertIn(key, DIMENSIONS)

    def test_each_dimension_has_four_facets(self):
        for key, info in DIMENSIONS.items():
            self.assertEqual(len(info["facets"]), 4, f"Dimension {key} should have 4 facets")

    def test_dimension_names(self):
        expected = {
            "D": "Discipline", "Y": "Yielding", "N": "Novelty", "A": "Acuity",
            "M": "Mercuriality", "I": "Impulsivity", "C": "Candour", "S": "Sociability",
        }
        for key, name in expected.items():
            self.assertEqual(DIMENSIONS[key]["name"], name)


class TestDynamicsProfile(unittest.TestCase):
    """Verify profile creation, validation, and serialisation."""

    def test_default_profile_is_valid(self):
        p = DynamicsProfile()
        self.assertTrue(p.validate())

    def test_out_of_range_fails_validation(self):
        p = DynamicsProfile(D=1.5)
        self.assertFalse(p.validate())

    def test_negative_fails_validation(self):
        p = DynamicsProfile(M=-0.1)
        self.assertFalse(p.validate())

    def test_json_round_trip(self):
        original = DynamicsProfile(D=0.71, Y=0.55, N=0.83, A=0.90, M=0.42, I=0.35, C=0.65, S=0.78)
        data = original.to_json()
        restored = DynamicsProfile.from_json(data)
        for dim in "DYNAMICS":
            self.assertAlmostEqual(getattr(original, dim), getattr(restored, dim), places=6)

    def test_json_keys_are_single_letters(self):
        p = DynamicsProfile()
        data = p.to_json()
        for key in "DYNAMICS":
            self.assertIn(key, data)

    def test_summary_contains_dimension_names(self):
        p = DynamicsProfile(D=0.9, S=0.1)
        summary = p.summary()
        self.assertIn("Discipline", summary)
        self.assertIn("Sociability", summary)

    def test_dimension_label_ranges(self):
        p = DynamicsProfile()
        self.assertEqual(p.dimension_label("D"), "moderate")  # default 0.5
        p2 = DynamicsProfile(D=0.05)
        self.assertEqual(p2.dimension_label("D"), "very low")
        p3 = DynamicsProfile(D=0.95)
        self.assertEqual(p3.dimension_label("D"), "very high")


class TestGeneration(unittest.TestCase):
    """Verify profile generation."""

    def test_generated_profile_is_valid(self):
        for _ in range(100):
            p = generate_profile()
            self.assertTrue(p.validate(), f"Generated profile failed validation: {p.to_json()}")

    def test_constraints_are_respected(self):
        p = generate_profile(constraints={"D": 0.8})
        self.assertAlmostEqual(p.D, 0.8, places=6)

    def test_octant_distribution(self):
        """Generate 200 profiles and check no single octant exceeds 30%."""
        octants = {}
        for _ in range(200):
            p = generate_profile()
            o = p.octant()
            octants[o] = octants.get(o, 0) + 1
        max_pct = max(octants.values()) / 200
        self.assertLess(max_pct, 0.30, "Single octant should not exceed 30% of generated profiles")


class TestCompatibility(unittest.TestCase):
    """Verify compatibility scoring."""

    def test_self_compatibility_is_high(self):
        p = generate_profile()
        score = compatibility_score(p, p)
        self.assertGreaterEqual(score, 0.5)

    def test_compatibility_range(self):
        for _ in range(50):
            a = generate_profile()
            b = generate_profile()
            score = compatibility_score(a, b)
            self.assertGreaterEqual(score, 0.0)
            self.assertLessEqual(score, 1.0)


class TestDerivations(unittest.TestCase):
    """Verify derivation functions return valid values."""

    def test_income_band_values(self):
        valid = {"low", "lower-middle", "middle", "upper-middle", "high"}
        for _ in range(50):
            p = generate_profile()
            self.assertIn(derive_income_band(p), valid)

    def test_spending_pattern_values(self):
        valid = {"frugal", "careful", "moderate", "generous", "impulsive"}
        for _ in range(50):
            p = generate_profile()
            self.assertIn(derive_spending_pattern(p), valid)

    def test_risk_tolerance_values(self):
        valid = {"very low", "low", "moderate", "high", "very high"}
        for _ in range(50):
            p = generate_profile()
            self.assertIn(derive_risk_tolerance(p), valid)

    def test_political_lean_range(self):
        for _ in range(50):
            p = generate_profile()
            lean = derive_political_lean(p)
            self.assertGreaterEqual(lean, -1.0)
            self.assertLessEqual(lean, 1.0)

    def test_high_discipline_high_novelty_tends_upper_income(self):
        """High D + high N should trend toward higher income bands."""
        upper_count = 0
        for _ in range(100):
            p = generate_profile(constraints={"D": 0.9, "N": 0.9})
            band = derive_income_band(p)
            if band in ("upper-middle", "high"):
                upper_count += 1
        self.assertGreater(upper_count, 50, "High D + N should produce upper income bands >50% of the time")


if __name__ == "__main__":
    unittest.main()
