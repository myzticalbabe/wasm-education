# Inclusion — Girls in STEM

This document records the design reasoning behind the inclusion features in
`demo-pages/kids-index.html`. It exists so that future contributors do not remove
these features as "decoration", and so that educators adopting the platform can see
what the claims rest on.

## The problem this addresses

Girls do not lose interest in STEM because of a drop in ability. The measured
pattern is a drop in **belonging**: interest peaks around age 11 and falls through
early adolescence, and the fall is larger for girls than for boys. Two mechanisms
are well documented:

1. **Stereotype threat** — when a task is framed as diagnostic of group ability,
   members of a negatively stereotyped group underperform relative to their own
   baseline.
2. **Role-model scarcity** — exposure to in-group experts in STEM measurably
   protects women's self-concept and anticipated success (Stout, Dasgupta,
   Hunsinger & McManus, 2011, *JPSP*), and counter-stereotypic role models reduce
   threat (Cheryan, Siy, Vichayapai, Drury & Kim, 2011, *Social Psychological and
   Personality Science*).

A gender-neutral page does nothing about either mechanism. Neutral is not neutral
in effect; it leaves the default signals in place. So the design goals below are
deliberate, not incidental.

## Design goals

| Goal | Implementation |
|---|---|
| Girls see women doing this exact work | **"Real Scientists" view** — working scientists whose actual contribution maps onto one of the four modules, with a "try this" tie-in |
| Effort is framed as the mechanism, not a consolation prize | Growth framing in copy: brains build connections by *using* them; the Brain Power Meter teaches that difficulty is the input to learning, not evidence against ability |
| No ranking surface | Badges mark exploration milestones, no leaderboard, no public comparison |
| Language is not gendered, but is not erased either | Content is inclusive of all learners while explicitly naming women scientists — including boys in the material without making girls the default audience for the soft version |
| Nothing asks a child to disclose identity | No accounts, no name field, no demographic prompt. Progress is a single local key. |

## The role models, and why these six

Each was chosen because her contribution *is* the module, so the connection can be
stated in one sentence a ten-year-old can follow. All facts below are checkable.

| Scientist | Contribution | Module |
|---|---|---|
| **Danielle Bassett** | Network neuroscientist; applies graph theory to how brain regions communicate | Word Web, Thought Paths |
| **Eve Marder** | Shows that a circuit's behaviour depends on the whole network, not on single neurons — "neurons don't fire alone" | Neuron Lab |
| **Rita Levi-Montalcini** | Discovered nerve growth factor; Nobel Prize in Physiology or Medicine, 1986 | Neuron Lab |
| **Brenda Milner** | Founded clinical neuropsychology; established that memory is several systems, not one | Word Web |
| **Radia Perlman** | Designed the spanning-tree algorithm that lets networks route without looping | Thought Paths |
| **Frances Allen** | First woman to win the Turing Award (2006); made compilers fast enough to matter | all four (why it runs at all) |

Deliberately excluded: "first woman to…" as the *headline* fact for anyone whose
work stands on its own. Where a first is mentioned (Allen, Turing Award), it is
paired with what she actually did.

## Facilitation guidance

The most common way these features get undermined is adult language around them:

- Say **"Bassett showed that brain networks reorganise when you learn"**, not
  "here's a female scientist".
- Say **"this is the part where it gets interesting"** when a learner is stuck,
  not "you're doing great" (person-praise is measurable *worse* than process-praise
  for persistence).
- Do not segregate by gender for pair work. The evidence supports mixed grouping
  with assigned, rotating roles.
- If a learner says "I'm bad at this", do not contradict it — ask what specifically
  felt hard, then point at it as a task feature with a named cause. That is the
  growth-mindset move that has evidence behind it; blanket reassurance does not.

## What is *not* claimed

The platform does not claim to close the gender gap. A single interactive page
cannot. What it claims is narrower and defensible: it removes two avoidable
contributors — absent representation and ability-framed feedback — from a resource
that otherwise has none.

## References

- Cheryan, S., Siy, J. O., Vichayapai, M., Drury, B. J., & Kim, S. (2011). Do female
  and male role models who embody STEM stereotypes hinder women's anticipated success
  in STEM? *Social Psychological and Personality Science*, 2(6), 656–664.
- Stout, J. G., Dasgupta, N., Hunsinger, M., & McManus, M. A. (2011). STEMing the
  tide: Using ingroup experts to inoculate women's self-concept in science,
  technology, engineering, and mathematics. *Journal of Personality and Social
  Psychology*, 100(2), 255–270.
- Dweck, C. S. (2006). *Mindset: The New Psychology of Success.* Random House.
- Mueller, C. M., & Dweck, C. S. (1998). Praise for intelligence can undermine
  children's motivation and performance. *JPSP*, 75(1), 33–52.
- AAUW (2010). *Why So Few? Women in Science, Technology, Engineering, and
  Mathematics.*
