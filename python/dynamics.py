# DYNAMICS-8 Personality Framework - Canonical Python Reference Implementation
# Copyright (c) 2026 Kronaxis Limited. Licensed under BSL 1.1. See LICENSE.
#
# Created by Jason Duke, Kronaxis Limited
# https://kronaxis.co.uk/dynamics
# Patent: UK Patent Application C (GB 2605150.8)

"""
Canonical Python library for the DYNAMICS-8 personality framework.

Eight continuous dimensions (0.0 to 1.0), each with four facets, designed
for behavioural prediction in digital environments. Zero external dependencies
(stdlib only: json, random, dataclasses, typing, math).

Dimension key:
    D = Discipline, Y = Yielding, N = Novelty, A = Acuity,
    M = Mercuriality, I = Impulsivity, C = Candour, S = Sociability

Specification: https://kronaxis.co.uk/dynamics
"""

from __future__ import annotations

import json
import math
import random
from dataclasses import dataclass, field
from typing import Any, Optional


# ---------------------------------------------------------------------------
# Dimension registry
# ---------------------------------------------------------------------------

DIMENSIONS: dict[str, dict[str, Any]] = {
    "D": {
        "name": "Discipline",
        "description": (
            "The degree to which a person is organised, diligent, and prudent "
            "versus careless, unstructured, and negligent."
        ),
        "facets": ["Organisation", "Diligence", "Perfectionism", "Prudence"],
    },
    "Y": {
        "name": "Yielding",
        "description": (
            "The degree to which a person yields to social pressure, tolerates "
            "disagreement, and adapts to persuasion versus resists influence."
        ),
        "facets": ["Patience", "Tolerance", "Flexibility", "Persuadability"],
    },
    "N": {
        "name": "Novelty",
        "description": (
            "The degree to which a person is intellectually inquisitive, "
            "aesthetically engaged, and unconventional versus incurious and "
            "conformist."
        ),
        "facets": ["Aesthetic sense", "Inquisitiveness", "Creativity", "Unconventionality"],
    },
    "A": {
        "name": "Acuity",
        "description": (
            "The degree to which a person is native to digital platforms, "
            "proficient with technology, privacy-aware, and comfortable across "
            "digital environments."
        ),
        "facets": ["Platform nativeness", "Content creation", "Privacy awareness", "Tech adoption"],
    },
    "M": {
        "name": "Mercuriality",
        "description": (
            "The degree to which a person is emotionally reactive, anxious, and "
            "temperamentally changeable versus stable, self-assured, and consistent."
        ),
        "facets": ["Anxiety", "Emotional reactivity", "Sentimentality", "Dependence"],
    },
    "I": {
        "name": "Impulsivity",
        "description": (
            "The degree to which a person acts on immediate urges without "
            "deliberation, seeks novel stimulation, and discounts future "
            "consequences."
        ),
        "facets": ["Delay discounting", "Sensation seeking", "Snap decision", "Boredom susceptibility"],
    },
    "C": {
        "name": "Candour",
        "description": (
            "The degree to which a person is transparent, fair, and modest "
            "versus manipulative, greedy, and self-aggrandising."
        ),
        "facets": ["Sincerity", "Fairness", "Modesty", "Materialism (inverse)"],
    },
    "S": {
        "name": "Sociability",
        "description": (
            "The degree to which a person is socially bold, energetic, and "
            "gregarious versus reserved, quiet, and solitary."
        ),
        "facets": ["Social boldness", "Liveliness", "Social esteem", "Gregariousness"],
    },
}

_DIM_KEYS: tuple[str, ...] = tuple(DIMENSIONS.keys())  # D, Y, N, A, M, I, C, S


# ---------------------------------------------------------------------------
# Band thresholds (score >= threshold -> label)
# ---------------------------------------------------------------------------

BANDS: list[tuple[float, str]] = [
    (0.80, "very high"),
    (0.60, "high"),
    (0.40, "moderate"),
    (0.20, "low"),
    (0.00, "very low"),
]


# ---------------------------------------------------------------------------
# Compatibility weights
# ---------------------------------------------------------------------------

COMPAT_WEIGHTS: dict[str, float] = {
    "D": 0.8,     # similar discipline levels reduce friction
    "Y": -0.3,    # mild complementarity: one yielding, one leading
    "N": 0.6,     # shared curiosity level aids connection
    "A": 0.5,     # similar digital fluency reduces communication gaps
    "M": -0.4,    # one stable partner balances one reactive partner
    "I": 0.2,     # slight similarity preference for pace alignment
    "C": 0.9,     # value alignment on honesty is critical
    "S": 0.4,     # moderate similarity in social energy
}


# ---------------------------------------------------------------------------
# Utility
# ---------------------------------------------------------------------------

def _clamp(val: float, lo: float = 0.0, hi: float = 1.0) -> float:
    """Clamp a value to the range [lo, hi]."""
    return max(lo, min(hi, val))


# ---------------------------------------------------------------------------
# Profile dataclass
# ---------------------------------------------------------------------------

@dataclass
class DynamicsProfile:
    """An eight-dimension DYNAMICS-8 personality profile.

    Each dimension is a float in [0.0, 1.0]. Optional per-dimension facets
    provide four sub-scores matching the facet order in the specification.

    Usage::

        p = DynamicsProfile(D=0.71, Y=0.55, N=0.82, A=0.45,
                            M=0.30, I=0.65, C=0.88, S=0.40)
        assert p.validate()
        print(p.summary())
        # D: Discipline = 0.71 (high), Y: Yielding = 0.55 (moderate), ...
        print(p.octant())
        # DH-YH-NH-AL-ML-IH-CH-SL
    """

    D: float = 0.5
    Y: float = 0.5
    N: float = 0.5
    A: float = 0.5
    M: float = 0.5
    I: float = 0.5
    C: float = 0.5
    S: float = 0.5
    facets: Optional[dict[str, list[float]]] = field(default=None, repr=False)

    # -- validation ----------------------------------------------------------

    def validate(self) -> bool:
        """Return True if all scores are within [0.0, 1.0] and structurally valid.

        Checks performed:
            - All eight dimension scores exist and are numeric.
            - All dimension scores lie in [0.0, 1.0].
            - If facets are provided, each facet list has exactly four entries,
              all within [0.0, 1.0], and keys are valid dimension letters.
        """
        for dim in _DIM_KEYS:
            val = getattr(self, dim, None)
            if val is None or not isinstance(val, (int, float)):
                return False
            if not (0.0 <= float(val) <= 1.0):
                return False
        if self.facets is not None:
            if not isinstance(self.facets, dict):
                return False
            for key, flist in self.facets.items():
                if key not in _DIM_KEYS:
                    return False
                if not isinstance(flist, list) or len(flist) != 4:
                    return False
                if not all(
                    isinstance(v, (int, float)) and 0.0 <= float(v) <= 1.0
                    for v in flist
                ):
                    return False
        return True

    # -- serialisation -------------------------------------------------------

    @classmethod
    def from_json(cls, data: dict) -> DynamicsProfile:
        """Parse a profile from the JSONB format used in soul_personas.dynamics.

        Accepts both core (8 keys) and extended (with ``facets`` sub-object)
        formats. Missing dimensions default to 0.5.

        Parameters
        ----------
        data : dict
            A dict with keys ``D``, ``Y``, ``N``, ``A``, ``M``, ``I``, ``C``,
            ``S`` (floats) and optionally ``facets`` (dict of dim -> [float]).
        """
        kwargs: dict[str, Any] = {}
        for dim in _DIM_KEYS:
            kwargs[dim] = float(data.get(dim, 0.5))
        raw_facets = data.get("facets")
        if raw_facets and isinstance(raw_facets, dict):
            facets: dict[str, list[float]] = {}
            for dim in _DIM_KEYS:
                if dim in raw_facets:
                    facets[dim] = [float(v) for v in raw_facets[dim]]
            kwargs["facets"] = facets if facets else None
        return cls(**kwargs)

    def to_json(self) -> dict:
        """Serialise to the JSONB format stored in soul_personas.dynamics.

        Dimension scores are rounded to four decimal places. Facets, if
        present, are included under a ``facets`` key.
        """
        out: dict[str, Any] = {
            dim: round(getattr(self, dim), 4) for dim in _DIM_KEYS
        }
        if self.facets:
            out["facets"] = {
                dim: [round(v, 4) for v in vals]
                for dim, vals in self.facets.items()
            }
        return out

    # -- band classification -------------------------------------------------

    @staticmethod
    def band(score: float) -> str:
        """Return the qualitative band label for a given score.

        Bands (inclusive lower bound, descending):
            [0.80, 1.00] -> "very high"
            [0.60, 0.80) -> "high"
            [0.40, 0.60) -> "moderate"
            [0.20, 0.40) -> "low"
            [0.00, 0.20) -> "very low"
        """
        for threshold, label in BANDS:
            if score >= threshold:
                return label
        return "very low"

    def dimension_label(self, dim: str) -> str:
        """Return the qualitative band label for a given dimension.

        Example: ``profile.dimension_label("D")`` -> ``"high"``
        """
        if dim not in _DIM_KEYS:
            raise KeyError(f"Unknown dimension: {dim}")
        return self.band(getattr(self, dim))

    # -- human-readable output -----------------------------------------------

    def summary(self) -> str:
        """Return a human-readable summary of the profile.

        Format::

            D: Discipline = 0.71 (high), Y: Yielding = 0.55 (moderate), ...
        """
        parts: list[str] = []
        for dim in _DIM_KEYS:
            val = getattr(self, dim)
            name = DIMENSIONS[dim]["name"]
            label = self.band(val)
            parts.append(f"{dim}: {name} = {val:.2f} ({label})")
        return ", ".join(parts)

    def octant(self) -> str:
        """Return a personality octant code based on a 0.5 threshold split.

        Each dimension is coded ``H`` (high, >= 0.5) or ``L`` (low, < 0.5).
        Result format: ``DH-YL-NH-AL-MH-IL-CH-SL`` (example).
        """
        codes: list[str] = []
        for dim in _DIM_KEYS:
            val = getattr(self, dim)
            codes.append(f"{dim}{'H' if val >= 0.5 else 'L'}")
        return "-".join(codes)

    # -- convenience access --------------------------------------------------

    def __getitem__(self, dim: str) -> float:
        """Allow dict-style access: ``profile['D']``."""
        if dim not in _DIM_KEYS:
            raise KeyError(f"Unknown dimension: {dim}")
        return getattr(self, dim)

    def as_dict(self) -> dict[str, float]:
        """Return the eight dimension scores as a plain dict."""
        return {dim: getattr(self, dim) for dim in _DIM_KEYS}


# ---------------------------------------------------------------------------
# Profile generation
# ---------------------------------------------------------------------------

def generate_profile(
    constraints: Optional[dict[str, Any]] = None,
) -> DynamicsProfile:
    """Generate a random DYNAMICS-8 profile.

    Constraints dict accepts three forms per dimension:

    - **Exact value:** ``{"D": 0.8}`` pins Discipline to exactly 0.8.
    - **Minimum bound:** ``{"S_min": 0.6}`` ensures Sociability >= 0.6.
    - **Maximum bound:** ``{"I_max": 0.3}`` ensures Impulsivity <= 0.3.

    Unconstrained dimensions are drawn from a beta(2, 2) distribution,
    which produces a bell-shaped spread centred on 0.5, matching realistic
    population norms. Constrained ranges use uniform sampling within bounds.

    Parameters
    ----------
    constraints : dict or None
        Optional constraints on dimension values. Keys are dimension letters
        (exact), or ``<dim>_min`` / ``<dim>_max`` for range bounds.

    Returns
    -------
    DynamicsProfile
        A validated profile with all scores in [0.0, 1.0].
    """
    constraints = constraints or {}
    scores: dict[str, float] = {}

    for dim in _DIM_KEYS:
        # Exact pin takes precedence.
        if dim in constraints:
            scores[dim] = _clamp(float(constraints[dim]))
            continue

        lo = float(constraints.get(f"{dim}_min", 0.0))
        hi = float(constraints.get(f"{dim}_max", 1.0))
        lo, hi = _clamp(lo), _clamp(hi)
        if lo > hi:
            lo, hi = hi, lo

        if lo == 0.0 and hi == 1.0:
            # Unconstrained: beta distribution for population-like spread.
            scores[dim] = _clamp(random.betavariate(2.0, 2.0))
        else:
            # Constrained range: uniform within bounds.
            scores[dim] = random.uniform(lo, hi)

    return DynamicsProfile(**scores)


# ---------------------------------------------------------------------------
# Compatibility
# ---------------------------------------------------------------------------

def compatibility_score(a: DynamicsProfile, b: DynamicsProfile) -> float:
    """Compute a 0.0-1.0 compatibility score between two profiles.

    Uses dimension-specific weights from ``COMPAT_WEIGHTS``:

    - **Positive weight:** similarity is rewarded (low difference = high score).
    - **Negative weight:** complementarity is rewarded (high difference = high
      score, i.e. "opposites attract").

    The final score is the weighted average across all eight dimensions,
    clamped to [0.0, 1.0].

    Parameters
    ----------
    a, b : DynamicsProfile
        The two profiles to compare.

    Returns
    -------
    float
        Compatibility score in [0.0, 1.0].
    """
    total = 0.0
    weight_sum = 0.0

    for dim in _DIM_KEYS:
        va, vb = a[dim], b[dim]
        w = COMPAT_WEIGHTS[dim]
        diff = abs(va - vb)

        if w >= 0:
            # Similarity: lower difference = higher contribution.
            dim_score = (1.0 - diff) * abs(w)
        else:
            # Complementarity: higher difference = higher contribution.
            dim_score = diff * abs(w)

        total += dim_score
        weight_sum += abs(w)

    return _clamp(total / weight_sum) if weight_sum > 0 else 0.5


# ---------------------------------------------------------------------------
# Derivation functions
#
# Each returns a raw float score. For income, spending, and risk the range
# is [0.0, 1.0]. For political lean it is [-1.0, +1.0].
#
# Callers can convert to a label using DynamicsProfile.band() where a
# categorical interpretation is needed.
# ---------------------------------------------------------------------------

_INCOME_BANDS = [
    (0.80, "high"),
    (0.60, "upper-middle"),
    (0.40, "middle"),
    (0.20, "lower-middle"),
    (0.00, "low"),
]

_SPENDING_LABELS = [
    (0.80, "frugal"),
    (0.60, "careful"),
    (0.40, "moderate"),
    (0.20, "generous"),
    (0.00, "impulsive"),
]

_RISK_LABELS = [
    (0.80, "very high"),
    (0.60, "high"),
    (0.40, "moderate"),
    (0.20, "low"),
    (0.00, "very low"),
]


def derive_income_band(p: DynamicsProfile) -> str:
    """Derive income-band label from Discipline and Novelty.

    Formula: ``D * 0.6 + N * 0.4``

    Rationale: self-regulation (D) and intellectual curiosity / career
    breadth (N) are the strongest personality predictors of earning
    potential (de Vries et al., 2009).

    Returns
    -------
    str
        One of: "low", "lower-middle", "middle", "upper-middle", "high".
    """
    score = _clamp(p.D * 0.6 + p.N * 0.4)
    for threshold, label in _INCOME_BANDS:
        if score >= threshold:
            return label
    return "low"


def derive_spending_pattern(p: DynamicsProfile) -> str:
    """Derive spending-pattern label from Candour, Discipline, and Impulsivity.

    Formula: ``C * 0.3 + D * 0.4 + (1 - I) * 0.3``

    High Candour (modesty, anti-materialism) and high Discipline (prudence)
    combined with low Impulsivity produce restrained spending. The inverse
    end of the scale indicates impulsive, unrestrained spending.

    Returns
    -------
    str
        One of: "impulsive", "generous", "moderate", "careful", "frugal".
    """
    score = _clamp(p.C * 0.3 + p.D * 0.4 + (1.0 - p.I) * 0.3)
    for threshold, label in _SPENDING_LABELS:
        if score >= threshold:
            return label
    return "impulsive"


def derive_risk_tolerance(p: DynamicsProfile) -> str:
    """Derive risk-tolerance label from Yielding, Impulsivity, and Mercuriality.

    Formula: ``Y * 0.25 + I * 0.40 + (1 - M) * 0.35``

    High Yielding (flexibility under pressure), high Impulsivity (sensation
    seeking, delay discounting), and low Mercuriality (emotional stability,
    absence of anxiety) produce the highest risk tolerance.

    Returns
    -------
    str
        One of: "very low", "low", "moderate", "high", "very high".
    """
    score = _clamp(p.Y * 0.25 + p.I * 0.40 + (1.0 - p.M) * 0.35)
    for threshold, label in _RISK_LABELS:
        if score >= threshold:
            return label
    return "very low"


def derive_political_lean(p: DynamicsProfile) -> float:
    """Derive political left-right lean from Novelty, Candour, and Yielding.

    Returns a float from -1.0 (far left) to +1.0 (far right).

    Mechanics:

    - **Left pull** (``N * 0.45``): high Novelty maps to openness to change,
      progressive values, and unconventionality.
    - **Right pull** (``(1 - C) * 0.35``): low Candour maps to materialism,
      status orientation, and competitive self-interest.
    - **Centre pull** (``Y * 0.20``): high Yielding dampens extremes via
      compliance, consensus-seeking, and social conformity.

    The raw difference (right_pull - left_pull) is compressed toward zero
    by the centre pull, modelling the tendency of agreeable individuals to
    avoid polarised positions.

    Parameters
    ----------
    p : DynamicsProfile
        The profile to evaluate.

    Returns
    -------
    float
        Score in [-1.0, 1.0]. Negative = left-leaning, positive = right-leaning.
    """
    left_pull = p.N * 0.45
    right_pull = (1.0 - p.C) * 0.35
    centre_pull = p.Y * 0.20

    # Raw lean: left_pull subtracts (toward -1), right_pull adds (toward +1).
    raw = right_pull - left_pull

    # High Yielding compresses the score toward zero.
    dampened = raw * (1.0 - centre_pull * 0.5)

    return _clamp(dampened, -1.0, 1.0)


# ---------------------------------------------------------------------------
# Module exports
# ---------------------------------------------------------------------------

__all__ = [
    "DIMENSIONS",
    "BANDS",
    "COMPAT_WEIGHTS",
    "DynamicsProfile",
    "generate_profile",
    "compatibility_score",
    "derive_income_band",
    "derive_spending_pattern",
    "derive_risk_tolerance",
    "derive_political_lean",
]
