# Batch Processing Kata: Ad-Fraud Detective

This is a practical kata to explore batch processing concepts from "Designing Data-Intensive Applications" Chapter 10. The challenge involves detecting click fraud patterns in large-scale clickstream data using distributed processing principles.

## The Problem

We need to identify two fraud patterns from one hour of click data:

- **Click Farms**: IPs with >100 clicks in one hour
- **Impossible Travel**: Users clicking from locations >500km apart within 10 minutes

## Key Observations from My Solution

I initially proposed reduce-side joins for everything, but learned this was inefficient. The geo dataset is small and static - perfect for map-side broadcast joins. This avoids shuffling the entire clicks dataset.

For the impossible travel detection, I recognized the reducer memory pitfall: loading all user clicks could cause O(n²) comparisons. The solution involves secondary sorting by timestamp before data reaches reducers.

## Implementation Rationale

The workflow should use separate job chains for each report. Click farms is straightforward: map by IP, count in reducer, filter results.

Impossible travel needs two stages: first enrich clicks with coordinates via broadcast join, then analyze user patterns with composite keys (user_id, timestamp) to enable efficient time-window comparisons in reducers.

## Questions for Further Exploration

When implementing the secondary sort for user click analysis, should we handle the 10-minute window in the reducer by loading adjacent time chunks, or is there a smarter partitioning strategy? Also, how would this solution change if we moved from classic MapReduce to Spark with in-memory DataFrames?

The distance calculation between cities - should this be pre-computed in the mapper during enrichment, or is it better to calculate on-demand in the reducer to avoid storing redundant coordinate pairs?
