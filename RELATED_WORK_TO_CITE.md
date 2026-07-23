# Related work to cite in the next DYNAMICS-8 revision

The published Zenodo record (DOI 10.5281/zenodo.19361059) is immutable; a citation added here is
folded in when a new version is cut. Each entry says what it is and why it belongs in DYNAMICS-8.

---

## Rahul & Chowdhury (2026) — persona panel Bayesian networks
- **Reference**: Kumar Rahul, Shovan Chowdhury. *Human–AI Construction of Bayesian Networks for
  Operational Decision Support: A Virtual Survey Approach.* arXiv:2607.14141 (2026). IIM Kozhikode.
- **Where it fits**: related work / applications of persona panels. It is an independent operations
  research application of the same core idea DYNAMICS-8 formalises — an LLM persona panel producing
  distributions over an outcome — used here to fill the conditional probability tables of a causal
  Bayesian network.
- **Why cite it**:
  1. It is adjacent prior art in a field (decision sciences / OR) our framework should reach, and it
     does not yet engage the silicon sampling literature, so citing it positions DYNAMICS-8 as the
     formal, validated basis it is missing.
  2. It is a clean contrast for our validation claim: they state plainly that their "most fundamental
     limitation is the absence of full real-world validation" and face validate only, whereas
     DYNAMICS-8 is validated against a published dataset with ten pre specified tests. Citing it lets us
     draw that line explicitly.
  3. Their interventional (do-calculus) framing is a natural "future directions" hook for extending
     DYNAMICS-8 panels from marginal prediction to causal, lever level reasoning.
- **Assessment**: `docs/kronaxis_intelligence/BBN_PERSONA_ELICITATION_ASSESSMENT_2026-07-23.md` (main repo).
