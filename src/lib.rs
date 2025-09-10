use serde::Deserialize;
use serde_repr::Deserialize_repr;

#[derive(Debug, Deserialize_repr, PartialEq)]
#[repr(u8)]
pub enum ResponseCode {
    Success = 0,
    NoResults = 1,
    InvalidParameter = 2,
    TokenNotFound = 3,
    TokenEmpty = 4,
    RateLimit = 5
}

#[derive(Debug, Deserialize, PartialEq)]
#[serde(rename_all = "lowercase")]
pub enum Difficulty {
    Easy,
    Medium,
    Hard
}

#[derive(Debug, Deserialize, PartialEq)]
#[serde(rename_all = "lowercase")]
pub enum AnswerType {
    Multiple,
    Boolean
}

#[derive(Debug, Deserialize)]
pub struct TriviaResponse {
    pub response_code: ResponseCode,
    pub results: Vec<TriviaQuestion>,
}

#[derive(Debug, Deserialize)]
pub struct TriviaQuestion {
    #[serde(rename = "type")]
    pub answer_type: AnswerType,
    pub difficulty: Difficulty,
    pub question: String,
    pub correct_answer: String,
    pub incorrect_answers: Vec<String>
}