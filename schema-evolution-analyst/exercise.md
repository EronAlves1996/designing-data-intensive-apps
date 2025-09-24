### Kata: Schema Evolution Analyst

**Objective:** Practice the critical skill of analyzing schema changes for backward and forward compatibility. You will not write code, but you will reason about the implications of schema modifications in a Protobuf/Avro-like context.

**Scenario: The User Service**
You are working on a service that stores user data. The initial schema for a `User` record is defined as follows:

```protobuf
// Schema Version 1.0
message User {
    required int32 user_id = 1;
    required string username = 2;
    optional string email_address = 3; // Newly added field
}
```

**Your Task (30 minutes):**

Analyze the following proposed schema changes. For each change, determine if it is:

1.  **Backward Compatible:** A new application (using the new schema) can read data written by an old application (using the old schema).
2.  **Forward Compatible:** An old application (using the old schema) can read data written by a new application (using the new schema).

State your compatibility verdict for each change and provide a brief (1-2 sentence) explanation for your reasoning, referencing concepts from the chapter.

---

**Proposed Changes to Analyze:**

1.  **Change A:** Add a new optional field `string full_name = 4;`.
    The schema will be:

```protobuf
message User {
  required int32 user_id = 1;
  required string username = 2;
  optional string email_address = 3;
  optional string full_name = 4;
}
```

**Backward Compatible:** Yes, the application will be backward compatible because the full_name now is optional. The old versions will not write the field and the new versions of app will not require the field to exists
**Forward Compatible:** Yes, because the new version can or cannot write the data, but, when the data is written, it's just ignored by old versions. 2. **Change B:** Change the field `username` from `required` to `optional`.
The schema will be:

```protobuf
message User {
  required int32 user_id = 1;
  optional string username = 2;
  optional string email_address = 3;
}
```

**Backward Compatible:** Yes, because old versions always will write username, making the field always present for new versions.
**Forward Compatible:** No, because new versions can omit the field, while old versions expect the field to be always present, breaking the compatibility. 3. **Change C:** Rename the field `username` to `display_name`.
The schema will be:

```protobuf
message User {
  required int32 user_id = 1;
  required string display_name = 2;
  optional string email_address = 3;
}
```

**Backward Compatible:** Yes. We can rename freely, the rule for compatibility in protobuf states that we cannot change is the tag index.
**Forward Compatible:** Yes. Same as before 4. **Change D:** Delete the field `email_address`.
The schema will be:

```protobuf
message User {
  required int32 user_id = 1;
  required string username = 2;
}
```

**Backward Compatible:** Yes. The field is always optional and we can delete freely, but we need to pay attention to not use the tag index again.
**Forward Compatible:** Yes, if tag index is not used again. > **Assessment notes:** If a new writer deletes the field and stops providing it, the old reader will receive a record where what it thinks is an optional field is simply absent. The application logic of the old reader must be able to handle a missing email_address gracefully. If the old application logic crashes when the field is null/missing, then forward compatibility is broken. So, while the serialization format may be forward compatible, the application might not be. > **My Answer:** But the application should expect what the serialization format mandates it to be. It should be likely an human fault to not align the application compatibility with the serialization compatibility, as the serialization format likely describe in some way the contract for the application. 5. **Change E:** Change the data type of `user_id` from `int32` to `int64`.
The schema will be:

```protobuf
message User {
  required int64 user_id = 1;
  required string username = 2;
  optional string email_address = 3;
}
```

**Backward Compatible:** Yes. By now on, we are expanding the size of the data type. No truncation is needed and no data is loose when new versions read.
**Forward Compatible:** No. If any old version read the user_id where it's a number where information only makes sense with the 32 upper bits, then we gonna have an information truncation and they will loose. 6. **Change F (Bonus - Real World Dilemma):** The product team wants to split the `email_address` field into two separate fields: `primary_email` (required) and `secondary_email` (optional). This is a fundamental change in the data model. How would you plan this migration in a way that minimizes service disruption? Outline the sequence of schema changes you would deploy.

1. The schema today have one field called `email_address`, and it accepts only one email. Rename `email_address` to `primary_email`.
2. Add the `secondary_email` schema as a field tag 4 and optional. New services can write it and old services can freely read discarding the value.

The two steps above can be made at one time or can be split. No need to do it in a multi-step way. > **Assessment notes:** A more robust, multi-step plan would be: > Add the new fields: Add primary_email (tag 4) and secondary_email (tag 5) as optional fields. Deploy the new schema. Update the application code to write to both the old email_address and the new primary_email field (a process called "dual writing"). This maintains full compatibility. > Migrate readers: Update all reader applications to first look for primary_email, and if not present, fall back to email_address. This makes readers resilient to both old and new data. > Migrate data: Once all applications are on the new version, run a batch job to backfill all existing records, copying the value from email_address to primary_email. > Remove the old field: After confirming that no application uses the email_address field anymore (after a sufficient grace period), remove it from the schema. This is the only backward-incompatible change, and it's done last when it's safe.
