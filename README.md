# DYNAMICS-8

An eight-dimension personality framework built for behavioural simulation.

DYNAMICS-8 extends the Big Five and HEXACO models with two dimensions purpose-built for digital and economic behaviour: **Acuity** (digital fluency) and **Impulsivity** (delay discounting and reward sensitivity). Each dimension is a continuous float from 0.0 to 1.0 with four granular facets, giving 32 behavioural parameters per persona.

| Code | Dimension | Lineage |
|------|-----------|---------|
| **D** | Discipline | Conscientiousness (Big Five) |
| **Y** | Yielding | Agreeableness (Big Five, HEXACO) |
| **N** | Novelty | Openness to Experience (Big Five) |
| **A** | Acuity | *New: digital fluency and platform nativeness* |
| **M** | Mercuriality | Neuroticism / Emotionality (Big Five, HEXACO) |
| **I** | Impulsivity | *New: delay discounting and reward sensitivity* |
| **C** | Candour | Honesty-Humility (HEXACO) |
| **S** | Sociability | Extraversion (Big Five, HEXACO) |

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
python -m pytest test_dynamics.py   # or: python -m unittest test_dynamics
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

- **Synthetic consumer panels**: census-weighted personas with realistic decision-making
- **AI training data**: personality-tagged causal reasoning traces
- **Market research**: stimulus testing across personality segments
- **Academic research**: population simulation for social science
- **Game AI**: personality-driven NPC behaviour
- **Chatbot personality**: consistent behavioural profiles for conversational agents

## Licence

- **Specification** ([DYNAMICS-8.md](DYNAMICS-8.md)): [Creative Commons Attribution 4.0 International](https://creativecommons.org/licenses/by/4.0/). Free to use, share, and adapt for any purpose including commercial use, with attribution.
- **Reference implementations** (`python/`, `go/`): [Business Source Licence 1.1](LICENSE-BSL). Free for non-production use. Commercial use requires a licence from Kronaxis. Converts to Apache 2.0 on 17 March 2030.

## Citation

```
Duke, J. (2026). DYNAMICS-8: A Purpose-Built Behavioural Simulation Framework.
Kronaxis Limited. https://kronaxis.co.uk/dynamics
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
