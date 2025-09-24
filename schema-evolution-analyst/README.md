# Schema Evolution Kata

This is a practical kata for studying Chapter 4 (Encoding and Evolution) of "Designing Data-Intensive Applications". The focus is on understanding schema compatibility in distributed systems without writing actual code.

## The Exercise

You're given an initial Protocol Buffers schema for a `User` record and need to analyze different schema changes for backward/forward compatibility. The goal is to reason about how each modification affects data flow between different service versions.

## What I Learned

The main insight is that schema evolution isn't just about the serialization format - it's about the contract between services. Protocol Buffers uses tag numbers, not field names, for identification, which makes renaming safe but type changes dangerous.

The tricky part was understanding that `optional` fields create an application-level contract. If the application logic doesn't handle missing optional fields properly, even technically compatible schema changes can break things. This distinction between serialization compatibility and application compatibility was the key learning.

For complex migrations like splitting fields, a multi-phase approach is safer than trying to do it in one shot. Dual writing and gradual reader migration are essential patterns.

## Implementation Rationale

The kata is designed to be solved through reasoning alone because in real systems, you'd use generated code from the schema. The important part is understanding the compatibility rules before touching any code.

For a real implementation, I'd focus on:

- Adding validation to ensure application logic properly handles optional fields
- Creating migration scripts that follow the safe evolution patterns identified
- Setting up schema registry checks to prevent incompatible changes

## Questions for Further Exploration

One thing I'm still thinking about: when you have multiple optional fields that are semantically related (like the email splitting scenario), how do you coordinate the migration across different services that might read the data at different times? Is there a pattern for handling these "semantic dependency" changes more cleanly?

Another question: in real Protobuf implementations, how do you actually enforce that old readers gracefully handle missing optional fields? Is this something you can catch at compile time or does it require runtime testing?

---

_Based on Chapter 4 of "Designing Data-Intensive Applications" by Martin Kleppmann_
