# UpdateTranscriptAttributionRequestBody


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `FromSpeaker`                                                      | **string*                                                          | :heavy_minus_sign:                                                 | The speaker to attribute the transcript to.                        |
| `FromUserID`                                                       | **string*                                                          | :heavy_minus_sign:                                                 | The user to attribute the transcript to.                           |
| `ToUserID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | The ID of the user to attribute the transcript to.                 |
| `TranscriptID`                                                     | **string*                                                          | :heavy_minus_sign:                                                 | The ID of the specific transcript entry to change attribution for. |
| `ConferenceBridgeID`                                               | **string*                                                          | :heavy_minus_sign:                                                 | The ID of the conference bridge to attribute the transcript to.    |