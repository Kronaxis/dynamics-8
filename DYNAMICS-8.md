# DYNAMICS-8 Personality Framework | Created by Jason Duke, Kronaxis Limited

**Version:** 1.0
**Author:** Jason Duke, Kronaxis Limited
**Date:** March 2026
**Licence:** Creative Commons Attribution 4.0 International (CC BY 4.0)
**Patent:** UK Patent Application C (GB 2605150.8)
**Website:** https://kronaxis.co.uk/dynamics
**Citation:** Duke, J. (2026). DYNAMICS-8: A Purpose-Built Behavioural Simulation Framework. Kronaxis Limited.

---

## 1. Abstract

DYNAMICS-8 is a continuous eight-dimension personality framework designed for behavioural simulation and computational prediction. It retains six dimensions grounded in thirty years of validated personality science (the Big Five and HEXACO models) and introduces two purpose-built dimensions: Acuity (digital fluency and platform nativeness) and Impulsivity (delay discounting and reward sensitivity), which close critical gaps in predicting observable digital and economic behaviour. Unlike the Big Five and HEXACO, which optimise for internal consistency of self-report questionnaires, DYNAMICS-8 optimises for predictive validity of observable behaviour in digital, financial, and social contexts. The framework produces deterministic derivations across these domains from a compact eight-value personality vector stored as JSONB.

---

## 2. The Eight Dimensions

Each dimension is scored on a continuous 0.0 to 1.0 scale. Each dimension has four constituent facets, also scored 0.0 to 1.0. The acronym DYNAMICS encodes the eight dimension initials: **D**iscipline, **Y**ielding, **N**ovelty, **A**cuity, **M**ercuriality, **I**mpulsivity, **C**andour, **S**ociability.

### 2.1. D: Discipline

**Full name:** Discipline
**Lineage:** Replaces Conscientiousness (Big Five, HEXACO)

The degree to which an individual is organised, diligent, and prudent versus careless, unstructured, and negligent. Discipline captures self-regulation, planning capacity, and routine adherence. It is the primary predictor of schedule consistency, financial reliability, and operational security compliance.

| Facet | Description |
|---|---|
| Organisation | Preference for order, planning, and structure |
| Diligence | Thoroughness in task completion and follow-through |
| Perfectionism | High standards and attention to detail |
| Prudence | Deliberation before acting; avoidance of unnecessary risk |

**Simulation effects:**

- **High D (> 0.65):** Regular schedules, consistent grammar and spelling, budgeted spending, app usage at fixed times, lower missed payment probability, higher OPSEC adherence, methodical navigation patterns.
- **Low D (< 0.35):** Erratic schedules, frequent typos, impulse purchases, irregular app usage patterns, higher financial default risk, inconsistent operational security, disorganised file and message management.

### 2.2. Y: Yielding

**Full name:** Yielding
**Lineage:** Replaces Agreeableness (Big Five, HEXACO)

The degree to which an individual yields to social pressure, tolerates disagreement, and adapts to persuasion versus resists influence, holds firm positions, and challenges others. Yielding reframes the traditional Agreeableness construct to emphasise susceptibility to persuasion and compliance with authority, which are directly measurable in advertising conversion, social proof response, and negotiation behaviour.

| Facet | Description |
|---|---|
| Patience | Tolerance for frustration and willingness to wait |
| Tolerance | Acceptance of different viewpoints without hostility |
| Flexibility | Willingness to compromise and adjust positions |
| Persuadability | Susceptibility to advertising, social proof, and authority arguments |

**Simulation effects:**

- **High Y (> 0.65):** Susceptible to targeted advertising, follows trends, changes opinions readily, compliant with authority, higher ad conversion rates, gentler voice tone, accommodating in group interactions.
- **Low Y (< 0.35):** Resistant to persuasion, fixed opinions, confrontational in discussions, lower ad conversion, requires evidence-based communication, dominates group decisions.

### 2.3. N: Novelty

**Full name:** Novelty
**Lineage:** Replaces Openness to Experience (Big Five, HEXACO)

The degree to which an individual is intellectually inquisitive, aesthetically engaged, and unconventional versus incurious, traditional, and conformist. Renamed from Openness to emphasise the curiosity and novelty-seeking core of the construct, which drives content consumption breadth, product adoption speed, and political orientation.

| Facet | Description |
|---|---|
| Aesthetic sense | Appreciation for art, beauty, and creative expression |
| Inquisitiveness | Drive to learn, explore, and understand |
| Creativity | Generation of novel ideas and unconventional solutions |
| Unconventionality | Receptiveness to non-traditional ideas, lifestyles, and values |

**Simulation effects:**

- **High N (> 0.65):** Diverse app usage, broad content consumption, high link-clicking rate, wide topic variety, left-leaning political tendency, early product adoption, varied social media interests.
- **Low N (< 0.35):** Narrow app usage, routine content consumption, conservative political tendency, strong brand loyalty, resistance to product switching, preference for familiar platforms.

### 2.4. A: Acuity

**Full name:** Acuity (Digital Acuity)
**Lineage:** New dimension. No equivalent in Big Five or HEXACO. Draws on digital literacy research (Hargittai, 2002; van Dijk, 2005).

The degree to which an individual is native to digital platforms, proficient with technology, privacy-aware, and comfortable across multiple digital environments. Acuity captures a behavioural axis that personality science has not previously isolated: the gap between having access to technology and being fluent in its use. This dimension is the primary predictor of phishing susceptibility, content creation quality, and cross-platform navigation speed.

| Facet | Description |
|---|---|
| Platform nativeness | Intuitive understanding of platform conventions, memes, and interaction norms |
| Content creation | Skill in producing digital content (text, image, video) appropriate to each platform |
| Privacy awareness | Understanding and use of privacy settings, VPNs, ad blockers, and data minimisation |
| Tech adoption | Speed of adopting new platforms, features, and technologies |

**Simulation effects:**

- **High A (> 0.65):** Multi-platform presence, platform-appropriate content, faster app navigation, privacy-conscious settings, high content creation quality, lower phishing susceptibility, early feature adoption.
- **Low A (< 0.35):** Single-platform usage, generic content, slower navigation, default privacy settings, vulnerability to phishing, low content creation frequency, late adoption of new features.

### 2.5. M: Mercuriality

**Full name:** Mercuriality
**Lineage:** Replaces Emotionality (HEXACO) and Neuroticism (Big Five, inverse polarity)

The degree to which an individual is emotionally reactive, anxious, and temperamentally changeable versus stable, self-assured, and consistent. Named for the mercurial temperament: rapid, unpredictable shifts in mood and emotional state. Aligned with HEXACO Emotionality polarity rather than the Big Five Neuroticism inverse, emphasising volatility over negativity.

| Facet | Description |
|---|---|
| Anxiety | Tendency to worry about future events and potential dangers |
| Emotional reactivity | Intensity of emotional responses to events and stimuli |
| Sentimentality | Strength of emotional bonds and empathic responses |
| Dependence | Need for emotional support and reassurance from others |

**Simulation effects:**

- **High M (> 0.65):** Rapid notification responses, emotionally charged messages, spending spikes after negative events, higher financial anxiety, warmer voice tone, crisis-driven social media activity, mood-dependent purchasing.
- **Low M (< 0.35):** Measured responses, consistent financial behaviour, lower social media activity during crises, stable spending patterns, cooler emotional register, steady engagement regardless of external events.

### 2.6. I: Impulsivity

**Full name:** Impulsivity
**Lineage:** New dimension. No direct equivalent in Big Five or HEXACO. Partial coverage by the UPPS Impulsive Behaviour Scale (Whiteside and Lynam, 2001) but not captured as an independent axis in either major personality model.

The degree to which an individual acts on immediate urges without deliberation, seeks novel stimulation, and discounts future consequences. While partially correlated with low Discipline (D), Impulsivity is a distinct behavioural driver. The correlation between D and I is moderate (r ~ -0.35) but leaves substantial independent variance. Impulsivity directly predicts digital interaction speed, purchase latency, doom-scrolling duration, and response to limited-time offers in ways that low Discipline alone does not.

| Facet | Description |
|---|---|
| Delay discounting | Preference for immediate rewards over larger delayed rewards |
| Sensation seeking | Drive toward novel, intense, and varied experiences |
| Snap decision | Tendency to act before fully evaluating options |
| Boredom susceptibility | Low tolerance for monotony; need for constant stimulation |

**Simulation effects:**

- **High I (> 0.65):** Instant notification responses, impulse purchases, rapid app-switching, doom-scrolling, burst typing, high susceptibility to limited-time offers, short session depth but high session count.
- **Low I (< 0.35):** Measured purchase decisions, delayed notification responses, focused single-app sessions, deliberate typing rhythm, resistance to urgency marketing, fewer but longer sessions.

### 2.7. C: Candour

**Full name:** Candour
**Lineage:** Replaces Honesty-Humility (HEXACO). Partial overlap with Agreeableness facets in Big Five.

The degree to which an individual is transparent, fair, and modest versus manipulative, greedy, and self-aggrandising. Direct replacement of HEXACO's sixth factor, renamed for clarity. Candour is the strongest single predictor of brand preference authenticity, social media curation intensity, and susceptibility to status-driven marketing.

| Facet | Description |
|---|---|
| Sincerity | Genuineness in social interaction; avoidance of flattery and manipulation |
| Fairness | Unwillingness to exploit others for personal gain |
| Modesty | Absence of self-importance and entitlement |
| Materialism (inverse) | Indifference to luxury, status symbols, and wealth display |

**Simulation effects:**

- **High C (> 0.65):** Authentic self-disclosure, modest spending, volunteer engagement, preference for brands signalling social responsibility, consistent cover narrative, lower susceptibility to aspirational advertising.
- **Low C (< 0.35):** Curated social media presence, self-promotional content, higher luxury spending, aspiration-driven consumption, status-seeking behaviour, responsive to prestige branding.

### 2.8. S: Sociability

**Full name:** Sociability
**Lineage:** Replaces Extraversion (Big Five) and eXtraversion (HEXACO)

The degree to which an individual is socially bold, energetic, and gregarious versus reserved, quiet, and solitary. Renamed from Extraversion to emphasise the observable social behaviour output rather than the internal energy construct. Sociability is the primary predictor of posting frequency, platform count, session count, and word-of-mouth amplification potential.

| Facet | Description |
|---|---|
| Social boldness | Confidence in social situations; willingness to initiate contact |
| Liveliness | Enthusiasm, energy, and optimism in interactions |
| Social esteem | Positive self-evaluation in social contexts |
| Gregariousness | Enjoyment of social gatherings and group activities |

**Simulation effects:**

- **High S (> 0.65):** Frequent posting, multiple platform activity, fast typing, high session counts, word-of-mouth amplification, higher initial follower potential, group activity participation.
- **Low S (< 0.35):** Lurker behaviour, minimal posting, brief sessions, single-platform preference, private purchasing, low social media output, preference for one-to-one communication.

---

## 3. Lineage

The following table maps DYNAMICS-8 dimensions to their ancestry in the Big Five (Costa and McCrae, 1992) and HEXACO (Lee and Ashton, 2004) models.

| DYNAMICS-8 | Code | Big Five Equivalent | HEXACO Equivalent | Relationship |
|---|---|---|---|---|
| Discipline | D | Conscientiousness | Conscientiousness (C) | Direct replacement. Identical construct, renamed for acronym clarity. |
| Yielding | Y | Agreeableness | Agreeableness (A) | Reframed. Emphasis shifted from cooperative disposition to persuadability and social compliance. |
| Novelty | N | Openness to Experience | Openness to Experience (O) | Direct replacement. Renamed to reflect the curiosity and novelty-seeking core. |
| Acuity | A | **No equivalent** | **No equivalent** | **New.** Captures digital platform fluency, privacy awareness, and technology adoption. |
| Mercuriality | M | Neuroticism (inverse polarity) | Emotionality (E) | Reframed. Aligned with HEXACO polarity. Named to emphasise volatility over negativity. |
| Impulsivity | I | **No equivalent** (partial: low Conscientiousness) | **No equivalent** (partial: low Conscientiousness) | **New.** Distinct from low Discipline. Captures delay discounting, sensation seeking, and boredom susceptibility as an independent axis. |
| Candour | C | No equivalent (partial: Agreeableness facet) | Honesty-Humility (H) | Direct replacement of HEXACO's sixth factor. Renamed for clarity and acronym fit. |
| Sociability | S | Extraversion | eXtraversion (X) | Direct replacement. Renamed to emphasise social behaviour output. |

**Summary:** Six of the eight DYNAMICS-8 dimensions are descendants of established constructs validated across 50+ cultures and 13+ languages. Two dimensions (Acuity and Impulsivity) are new, purpose-built for digital and economic behaviour prediction.

---

## 4. Scoring

### 4.1. Scale

All dimensions and facets use a continuous floating-point scale from 0.0 (minimum expression) to 1.0 (maximum expression). The midpoint of 0.5 represents the population mean. Values are not ordinal categories; the distance between 0.3 and 0.4 is identical to the distance between 0.7 and 0.8.

### 4.2. Band Labels

For qualitative reporting, scores map to five bands:

| Band | Range |
|---|---|
| Very Low | 0.00 to 0.19 |
| Low | 0.20 to 0.39 |
| Moderate | 0.40 to 0.59 |
| High | 0.60 to 0.79 |
| Very High | 0.80 to 1.00 |

Boundary conditions: 0.20 is Low (not Very Low). 0.40 is Moderate (not Low). 0.60 is High (not Moderate). 0.80 is Very High (not High).

### 4.3. Core Format

The core format stores eight top-level dimension scores. This is the minimum viable representation of a DYNAMICS-8 profile.

```
{D, Y, N, A, M, I, C, S}  where each value is in [0.0, 1.0]
```

### 4.4. Extended Format

The extended format adds four facet scores per dimension (32 facets total). Facet order within each dimension is fixed and matches the order specified in Section 2.

```
{D, Y, N, A, M, I, C, S, facets: {D: [f1, f2, f3, f4], ..., S: [f1, f2, f3, f4]}}
```

Where each `f1...f4` is in [0.0, 1.0].

### 4.5. Reconstruction from Core

When only core scores are available, facets may be estimated by applying Gaussian noise (sigma = 0.05) around the dimension mean, clamped to [0.0, 1.0]. This is a lossy reconstruction; extended format should be used when facet-level precision matters.

---

## 5. Storage Schema

DYNAMICS-8 profiles are stored as JSONB (binary JSON) in relational databases. The canonical storage column is `soul_personas.dynamics` typed as `JSONB NOT NULL` in PostgreSQL. Two formats are defined.

### 5.1. Core Schema (8 scores)

```json
{
  "D": 0.71,
  "Y": 0.55,
  "N": 0.83,
  "A": 0.90,
  "M": 0.42,
  "I": 0.35,
  "C": 0.65,
  "S": 0.78
}
```

All eight keys are mandatory. Values must be numeric floats in the range [0.0, 1.0]. Unknown or unscored dimensions default to 0.5.

### 5.2. Extended Schema (8 scores + 32 facets)

```json
{
  "D": 0.71,
  "Y": 0.55,
  "N": 0.83,
  "A": 0.90,
  "M": 0.42,
  "I": 0.35,
  "C": 0.65,
  "S": 0.78,
  "facets": {
    "D": [0.75, 0.68, 0.72, 0.69],
    "Y": [0.50, 0.60, 0.55, 0.55],
    "N": [0.85, 0.80, 0.88, 0.79],
    "A": [0.92, 0.88, 0.90, 0.90],
    "M": [0.50, 0.38, 0.45, 0.35],
    "I": [0.30, 0.40, 0.35, 0.35],
    "C": [0.70, 0.60, 0.65, 0.65],
    "S": [0.80, 0.75, 0.82, 0.75]
  }
}
```

The `facets` key is optional. When present, each dimension key maps to an array of exactly four floats in fixed order matching the facet definitions in Section 2. All facet values must be in [0.0, 1.0].

**Facet index reference:**

| Dimension | Index 0 | Index 1 | Index 2 | Index 3 |
|---|---|---|---|---|
| D | Organisation | Diligence | Perfectionism | Prudence |
| Y | Patience | Tolerance | Flexibility | Persuadability |
| N | Aesthetic sense | Inquisitiveness | Creativity | Unconventionality |
| A | Platform nativeness | Content creation | Privacy awareness | Tech adoption |
| M | Anxiety | Emotional reactivity | Sentimentality | Dependence |
| I | Delay discounting | Sensation seeking | Snap decision | Boredom susceptibility |
| C | Sincerity | Fairness | Modesty | Materialism (inverse) |
| S | Social boldness | Liveliness | Social esteem | Gregariousness |

---

## 6. Derivation Mappings

DYNAMICS-8 dimensions drive deterministic behavioural derivations. Each derived attribute is a weighted function of one or more dimension scores. Where a dimension is marked "(inverse)", the derivation uses `1.0 - score`.

### 6.1. Financial Derivations

| Derived Attribute | Formula | Description |
|---|---|---|
| Income band | `D * 0.6 + N * 0.4` | Discipline and intellectual breadth predict earning capacity. Bands: low (< 0.25), lower-middle (0.25-0.40), middle (0.40-0.60), upper-middle (0.60-0.75), high (>= 0.75). |
| Spending pattern | `C * 0.3 + D * 0.4 + (1-I) * 0.3` | Candour moderates materialism; Discipline constrains spending; inverse Impulsivity rewards restraint. Higher score = more restrained. Bands: impulsive (< 0.25), generous (0.25-0.40), moderate (0.40-0.60), careful (0.60-0.75), frugal (>= 0.75). |
| Risk tolerance | `Y * 0.25 + I * 0.40 + (1-M) * 0.35` | Yielding permits risk; Impulsivity embraces it; emotional stability (inverse M) enables it. Bands: very low (< 0.25), low (0.25-0.40), moderate (0.40-0.60), high (0.60-0.75), very high (>= 0.75). |
| Financial anxiety | `M * 0.7 + (1 - min(balance/income, 1.0)) * 0.3` | Mercuriality amplifies emotional reactivity to financial pressure. Context-dependent: requires balance and income as external inputs. |
| Credit score proxy | `D * 0.5 + (1-I) * 0.3 + (1-M) * 0.2` | Discipline builds credit history; low Impulsivity avoids missed payments; emotional stability prevents panic decisions. |

### 6.2. Political and Social Derivations

| Derived Attribute | Formula | Description |
|---|---|---|
| Political lean | `left_pull = N * 0.45; right_pull = (1-C) * 0.35; raw = right_pull - left_pull; dampened = raw * (1 - Y * 0.20 * 0.5)` | Returns float from -1.0 (far left) to 1.0 (far right). High Novelty pulls left (progressive values). Low Candour pulls right (materialism, status orientation). High Yielding dampens toward centre (consensus-seeking). |
| Social capital | `S * 0.6 + Y * 0.4` | Sociability builds network breadth; Yielding builds cooperative ties. |

### 6.3. Digital Behaviour Derivations

| Derived Attribute | Primary Inputs | Description |
|---|---|---|
| Keystroke timing | S, D, I, A | Social energy and impulsivity accelerate typing speed; discipline and digital fluency modulate rhythm and error correction. |
| Keystroke error rate | M, (1-D), (1-A) | Emotional reactivity, low discipline, and low digital fluency increase error frequency. |
| Notification response latency | I, M, S | Impulsivity and emotional reactivity drive urgency; sociability amplifies response speed for social notifications. |
| Doom-scrolling duration | I, N, S | Impulsivity sustains scrolling; novelty-seeking widens content appetite; sociability adds social content pull. |
| OPSEC score | D, A, C | Discipline ensures process adherence; digital fluency enables tool use; candour prevents careless disclosure. |
| Session frequency | S, I, A | Social energy and impulsivity increase session count; digital fluency lowers friction to re-engagement. |
| Content creation quality | A, N, D | Platform fluency, creativity, and attention to detail determine output quality. |

### 6.4. Voice and Communication Derivations

| Derived Attribute | Primary Inputs | Description |
|---|---|---|
| Voice warmth | M, Y | Mercuriality adds emotional colour; Yielding adds gentleness. |
| Voice pace | S, I | Sociability and impulsivity increase speaking rate. |
| Voice precision | D, A | Discipline and digital fluency produce articulate, structured speech. |

### 6.5. Consumer Behaviour Derivations

| Derived Attribute | Primary Inputs | Description |
|---|---|---|
| Brand loyalty | (1-N), D, (1-I) | Low novelty-seeking, high discipline, and low impulsivity produce strong brand attachment. |
| Ad conversion rate | Y, I, (1-A) | Persuadability, impulsivity, and low privacy awareness increase conversion. |
| Content sharing propensity | S, N, (1-C) | Social energy, curiosity, and self-promotion drive sharing. |

---

## 7. Interaction Rules

### 7.1. Compatibility Scoring

Compatibility between two DYNAMICS-8 profiles is computed as a weighted function across all eight dimensions. Each dimension carries a weight with a sign: positive weights reward similarity (low difference between profiles); negative weights reward complementarity (high difference between profiles).

| Dimension | Weight | Direction | Rationale |
|---|---|---|---|
| D (Discipline) | +0.8 | Similarity | Similar discipline levels reduce friction in shared routines and planning. |
| Y (Yielding) | -0.3 | Complementarity | One yielding partner and one leading partner produces stable role differentiation. |
| N (Novelty) | +0.6 | Similarity | Shared curiosity level aids mutual engagement and reduces boredom asymmetry. |
| A (Acuity) | +0.5 | Similarity | Similar digital fluency reduces communication gaps and platform misalignment. |
| M (Mercuriality) | -0.4 | Complementarity | One stable partner balances one reactive partner, providing emotional regulation. |
| I (Impulsivity) | +0.2 | Similarity | Slight similarity preference for pace alignment in decision-making. |
| C (Candour) | +0.9 | Similarity | Value alignment on honesty and transparency is the strongest predictor of trust. |
| S (Sociability) | +0.4 | Similarity | Moderate similarity in social energy reduces withdrawal/overwhelm conflicts. |

### 7.2. Computation

For each dimension `d` with weight `w`, given profiles `a` and `b`:

```
diff = abs(a[d] - b[d])
if w >= 0:
    dim_score = (1.0 - diff) * abs(w)    # similarity: lower difference = higher score
else:
    dim_score = diff * abs(w)              # complementarity: higher difference = higher score

compatibility = sum(dim_scores) / sum(abs(weights))
```

Result is clamped to [0.0, 1.0].

---

## 8. Profile Generation

### 8.1. Unconstrained Generation

Unconstrained dimensions are drawn from a Beta distribution with parameters alpha = 2, beta = 2. This produces a unimodal distribution centred at 0.5 with moderate variance (mean = 0.5, variance ~ 0.05), which matches the expected population distribution of personality traits more closely than a uniform distribution.

### 8.2. Constrained Generation

Constraints are specified as a dictionary with three key types:

| Constraint Type | Format | Example | Effect |
|---|---|---|---|
| Exact | `{"D": 0.8}` | Pin Discipline to 0.8 | Dimension set to the exact value. |
| Minimum | `{"S_min": 0.6}` | Sociability >= 0.6 | Uniform sampling within [0.6, 1.0]. |
| Maximum | `{"I_max": 0.3}` | Impulsivity <= 0.3 | Uniform sampling within [0.0, 0.3]. |
| Range | `{"N_min": 0.4, "N_max": 0.7}` | Novelty in [0.4, 0.7] | Uniform sampling within [0.4, 0.7]. |

Constrained ranges use uniform sampling (not Beta) because the constraint itself shapes the distribution and a Beta distortion within a narrow range would produce unintended bias.

---

## 9. Validation Rules

### 9.1. Score Validation

1. All eight dimension keys (D, Y, N, A, M, I, C, S) must be present at the top level.
2. All dimension values must be numeric and within [0.0, 1.0].
3. No additional keys are permitted at the top level beyond the eight dimension codes and `facets`.

### 9.2. Facet Validation

4. If `facets` is present, it must contain all eight dimension keys.
5. Each facet array must contain exactly four numeric values within [0.0, 1.0].
6. Facet order is fixed per the table in Section 5.2 and must not be reordered.

### 9.3. Population Distribution Validation (Octant Rule)

7. When generating a population of profiles, no single personality octant may contain more than 20% of the population. An octant is defined by binarising each dimension at the 0.5 midpoint (High if >= 0.5, Low if < 0.5), producing 2^8 = 256 possible octants. This rule prevents degenerate clustering in generated datasets.

### 9.4. Internal Consistency

8. If both core scores and facets are present, the mean of the four facets for a dimension should fall within 0.15 of the core dimension score. This is a soft constraint (warning, not rejection) to accommodate intentional facet dispersion.

---

## 10. Comparison Table

| Property | DYNAMICS-8 | Big Five (NEO PI-R) | HEXACO-PI-R |
|---|---|---|---|
| **Dimensions** | 8 | 5 | 6 |
| **Design purpose** | Behavioural simulation and computational prediction | Self-report personality measurement | Self-report personality measurement |
| **Validation criterion** | Predictive validity of observable digital and economic behaviour | Internal consistency of self-report items (Cronbach's alpha) | Internal consistency of self-report items (Cronbach's alpha) |
| **Digital behaviour prediction** | Native (Acuity dimension + derivation mappings) | Not addressed | Not addressed |
| **Impulse prediction** | Native (Impulsivity dimension) | Partial (low Conscientiousness facet) | Partial (low Conscientiousness facet) |
| **Honesty/ethics dimension** | Yes (Candour) | No (scattered across Agreeableness facets) | Yes (Honesty-Humility) |
| **Facets per dimension** | 4 (32 total) | 6 (30 total) | 4 (24 total) |
| **Scoring range** | Continuous 0.0 to 1.0 | Likert 1-5 or T-scores | Likert 1-5 |
| **Questionnaire required** | No (generative; can be inferred from observed behaviour) | Yes (240 items, NEO PI-R) | Yes (200 items, HEXACO-PI-R) |
| **Storage format** | JSONB (8 or 40 floats) | Not specified | Not specified |
| **Cross-cultural validation** | Inherits from Big Five/HEXACO (50+ cultures) for 6 dimensions; 2 new dimensions grounded in UPPS and digital literacy literature | 50+ cultures | 50+ cultures |

---

## 11. Dimension Interactions

Dimensions do not operate in isolation. The following interaction patterns produce emergent behavioural signatures that exceed what any single dimension predicts.

| Pattern | Dimensions | Effect |
|---|---|---|
| Financial chaos | High I + Low D | Impulse purchases, missed payments, boom-bust spending cycles. Strongest predictor of consumer debt accumulation. |
| Early tech adoption | High N + High A | First to trial new platforms, first to adopt features, highest-quality content on emerging platforms. |
| Persuasion vulnerability | High Y + High I + Low D | Maximum susceptibility to advertising, social proof, and limited-time offers. Highest-value segment for direct-response marketing. |
| Digital fortification | High A + High D + High C | Privacy-conscious, strong passwords, ad blockers, minimal digital footprint. Highest OPSEC scores. Lowest phishing susceptibility. |
| Social amplifier | High S + High N + Low C | Prolific, aspirational content creators. Effective brand ambassadors. May overstate product benefits. |
| Emotional digital volatility | High M + High I + High S | Rapid notification responses, emotionally charged public posts, doom-scrolling during crises, mood-triggered purchasing. |
| Analytical sceptic | Low Y + High N + High A | Thorough researchers, reject authority claims, require empirical evidence. Lowest ad click-through rate but highest post-purchase satisfaction. |
| Stable loyalist | High D + Low N + Low I | Most brand-loyal consumer segment. Routine adherence, resistance to competitor switching. High retention, high acquisition cost. |

---

## 12. Citation Format

### Academic Citation

> Duke, J. (2026). DYNAMICS-8: A Purpose-Built Behavioural Simulation Framework. Kronaxis Limited. Available at: https://kronaxis.co.uk/dynamics

### BibTeX

```bibtex
@techreport{duke2026dynamics8,
  author       = {Duke, Jason},
  title        = {{DYNAMICS-8}: A Purpose-Built Behavioural Simulation Framework},
  institution  = {Kronaxis Limited},
  year         = {2026},
  url          = {https://kronaxis.co.uk/dynamics},
  note         = {Version 1.0. CC BY 4.0.}
}
```

### In-Text Reference

When referencing DYNAMICS-8 dimensions in running text, use the full dimension name followed by the letter code in parentheses on first use: "Discipline (D)". Subsequent references may use the letter code alone.

When referencing the framework itself: "the DYNAMICS-8 framework (Duke, 2026)".

---

## 13. Licence

### Specification Licence

This specification is released under the **Creative Commons Attribution 4.0 International (CC BY 4.0)** licence. You are free to share and adapt this material for any purpose, including commercial use, provided you give appropriate credit.

Full licence text: https://creativecommons.org/licenses/by/4.0/

### Patent Notice

The computational methods described in the Derivation Mappings (Section 6), Interaction Rules (Section 7), and their application to behavioural simulation are the subject of UK Patent Application GB 2605150.8 (Consumer Behaviour Simulation), filed 10 March 2026. The CC BY 4.0 licence covers the specification text and the dimensional framework definitions. It does not grant a licence to any patent claims. Commercial implementations of the derivation formulas and interaction algorithms should seek independent legal advice regarding patent clearance.

### Reference Implementations

Reference implementations of the DYNAMICS-8 framework in Python and Go are available under the Business Source Licence 1.1 (BSL 1.1) as part of the Kronaxis Forge and Kronaxis Panel Studio open-core distributions. The BSL 1.1 permits non-production use and converts to Apache 2.0 after the change date.

### Contact

Kronaxis Limited
Website: https://kronaxis.co.uk
DYNAMICS-8 information: https://kronaxis.co.uk/dynamics

---

## References

Costa, P. T., and McCrae, R. R. (1992). *Revised NEO Personality Inventory (NEO PI-R) and NEO Five-Factor Inventory (NEO-FFI) Professional Manual.* Odessa, FL: Psychological Assessment Resources.

Goldberg, L. R. (1990). An alternative "description of personality": The Big-Five factor structure. *Journal of Personality and Social Psychology,* 59(6), 1216-1229.

Gray, J. A. (1970). The psychophysiological basis of introversion-extraversion. *Behaviour Research and Therapy,* 8(3), 249-266.

Hargittai, E. (2002). Second-level digital divide: Differences in people's online skills. *First Monday,* 7(4).

Lee, K., and Ashton, M. C. (2004). Psychometric properties of the HEXACO Personality Inventory. *Multivariate Behavioral Research,* 39(2), 329-358.

van Dijk, J. A. G. M. (2005). *The Deepening Divide: Inequality in the Information Society.* Thousand Oaks, CA: Sage.

Whiteside, S. P., and Lynam, D. R. (2001). The Five Factor Model and impulsivity: Using a structural model of personality to understand impulsivity. *Personality and Individual Differences,* 30(4), 669-689.
