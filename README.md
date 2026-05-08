# DYNAMICS-8

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-blue.svg)](https://creativecommons.org/licenses/by/4.0/)

> **Two new personality dimensions for the digital age — built for behavioural simulation, not for self-help.** DYNAMICS-8 takes the validated Big Five and HEXACO frameworks and adds **Acuity** (platform fluency, digital nativeness) and **Impulsivity** (delay discounting, reward sensitivity), the two axes that drive online decision-making but aren't captured by personality science designed in the 1990s.

The result: a continuous 0.0–1.0 score on each of eight dimensions, four facets per dimension (32 parameters total), deterministically derivable into economic behaviour, political orientation, and population-segmentation cells. Used in production by [Panel Studio](https://github.com/Kronaxis/kronaxis-panel-studio) to simulate 500–65,000 personas at a time, and validated publicly by [KPM-1](https://github.com/Kronaxis/kpm1-election-projections) (pre-registered, hash-committed UK election predictions).

## How it differs from Big Five and HEXACO

| Framework | Dimensions | Built for | Captures digital fluency? | Captures impulsivity / delay discounting? |
|---|---|---|---|---|
| **Big Five (OCEAN)** | 5 | General personality science | ✗ | Partially (under Conscientiousness) |
| **HEXACO** | 6 | Adds honesty/humility | ✗ | ✗ |
| **DYNAMICS-8** | 8 | Behavioural simulation in the digital economy | ✓ (Acuity) | ✓ (Impulsivity) |

Six of DYNAMICS-8's dimensions map cleanly to Big Five / HEXACO so you can interoperate with existing literature. The two new ones (A and I) are the load-bearing additions that make synthetic personas behave like real people online — not just on personality tests.

| Code | Dimension | Lineage |
|------|-----------|---------|
| **D** | Discipline | Conscientiousness (Big Five) |
| **Y** | Yielding | Agreeableness (Big Five, HEXACO) |
| **N** | Novelty | Openness to Experience (Big Five) |
| **A** | Acuity | **New: digital fluency and platform nativeness** |
| **M** | Mercuriality | Neuroticism / Emotionality (Big Five, HEXACO) |
| **I** | Impulsivity | **New: delay discounting and reward sensitivity** |
| **C** | Candour | Honesty-Humility (HEXACO) |
| **S** | Sociability | Extraversion (Big Five, HEXACO) |

## Part of the Kronaxis research stack

DYNAMICS-8 is the foundation of a four-project open-source stack:

1. **DYNAMICS-8** (this repo) — the framework and reference implementations
2. [**Panel Studio**](https://github.com/Kronaxis/kronaxis-panel-studio) — the engine that simulates 500–65,000 DYNAMICS-tagged personas at a time
3. [**KPM-1**](https://github.com/Kronaxis/kpm1-election-projections) — pre-registered, hash-verified election predictions (the public proof the stack works)
4. [**Kronaxis Router**](https://github.com/Kronaxis/kronaxis-router) — the LLM proxy that makes running 65,000 simulated personas economically viable

Each piece is independently usable; together they cover the loop from psychographic framework → simulated population → public falsifiable forecast → cost-efficient inference at scale.

## What it does

From an eight-value personality vector, DYNAMICS-8 deterministically derives:

- **Economic behaviour:** income band, spending pattern, risk tolerance
- **Political orientation:** left-right lean with centre-pull dampening
- **Compatibility:** weighted similarity/complementarity scoring between two profiles
- **Band classification:** five-level qualitative labels per dimension
- **Octant coding:** 256-cell personality space for population segmentation

## Quick start

### Python

```python
from dynamics import DynamicsProfile, generate_profile, derive_income_band

# Generate a random profile (Beta(2,2) distribution)
profile = generate_profile()
print(profile.summary())
# D: Discipline = 0.71 (high), Y: Yielding = 0.55 (moderate), ...

# Create a specific profile
p = DynamicsProfile(D=0.82, Y=0.45, N=0.91, A=0.73, M=0.28, I=0.35, C=0.88, S=0.62)
print(derive_income_band(p))  # "high"
print(p.octant())             # "DH-YL-NH-AH-ML-IL-CH-SH"

# Constrained generation
researcher = generate_profile(constraints={"N": 0.9, "A": 0.85, "D_min": 0.6})
```

Zero external dependencies. Copy `python/dynamics.py` into your project and import it.

```bash
cd python
python -m unittest test_dynamics    # zero dependencies
python -m pytest test_dynamics.py   # alternative, if pytest is installed
```

### Go

```go
import "github.com/kronaxis/dynamics-8/go"

p := dynamics.GenerateProfile()
fmt.Println(p.Summary())
fmt.Println(dynamics.DimensionLabel(p.D))  // "high"
fmt.Printf("Income: %.2f\n", dynamics.DeriveIncomeBand(p))
fmt.Printf("Compatibility: %.2f\n", dynamics.CompatibilityScore(p, other))
```

```bash
go get github.com/kronaxis/dynamics-8/go
```

## Specification

The full formal specification is in [DYNAMICS-8.md](DYNAMICS-8.md). It covers:

- All 8 dimensions with facet definitions and simulation effects
- Scoring bands and thresholds
- Derivation formulas (income, spending, risk, political lean)
- Compatibility algorithm with per-dimension weights
- Validation rules (9 rules including octant distribution and internal consistency)
- Storage schema (JSONB format)
- Comparison table: Big Five vs HEXACO vs DYNAMICS-8

Read it online at [kronaxis.co.uk/dynamics](https://kronaxis.co.uk/dynamics).

## Use cases

- **Synthetic consumer panels**: census weighted personas with realistic decisions
- **AI training data**: personality tagged causal reasoning traces
- **Market research**: stimulus testing across personality segments
- **Academic research**: population simulation for social science
- **Game AI**: NPC behaviour driven by personality
- **Chatbot personality**: consistent behavioural profiles for conversational agents

## Licence

- **Specification** ([DYNAMICS-8.md](DYNAMICS-8.md)): [Creative Commons Attribution 4.0 International](https://creativecommons.org/licenses/by/4.0/). Free to use, share, and adapt for any purpose including commercial use, with attribution.
- **Reference implementations** (`python/`, `go/`): [Business Source Licence 1.1](LICENSE-BSL). Free for non-production use. Commercial use requires a licence from Kronaxis. Converts to Apache 2.0 on 17 March 2030.

## Citation

```
Duke, J. (2026). DYNAMICS-8: An Eight-Dimension Personality Framework for
Computational Behavioural Simulation. Zenodo. https://doi.org/10.5281/zenodo.19361059
```

```bibtex
@article{duke_dynamics8_2026,
  title={DYNAMICS-8: An Eight-Dimension Personality Framework for Computational Behavioural Simulation},
  author={Duke, Jason},
  year={2026},
  doi={10.5281/zenodo.19361059},
  url={https://doi.org/10.5281/zenodo.19361059},
  publisher={Zenodo}
}
```

## Patent notice

The derivation mappings and interaction algorithms described in the specification are the subject of UK Patent Application GB 2605150.8 (Consumer Behaviour Simulation), filed 10 March 2026. The CC BY 4.0 licence covers the specification text and dimensional framework definitions. It does not grant a licence to patent claims. See [DYNAMICS-8.md](DYNAMICS-8.md) Section 10 for details.

## Links

- [Full specification](https://kronaxis.co.uk/dynamics)
- [Kronaxis Panel Studio](https://github.com/kronaxis/kronaxis-panel-studio) (synthetic consumer panel platform using DYNAMICS-8)
- [HuggingFace dataset](https://huggingface.co/datasets/kronaxis/uk-personas-500) (500 DYNAMICS-8 profiled personas)
- [kronaxis.co.uk](https://kronaxis.co.uk)

---

Created by Jason Duke, [Kronaxis Limited](https://kronaxis.co.uk). Built in England.
