use serde::Deserialize;

#[derive(Deserialize, Clone)]
pub struct CreateComplaintDTO {
    pub user_id: i32,
    pub user_name: String,
    pub subject_name: String,
    pub complaint_subject_id: i32,
    pub complaint_text: String,
}

#[derive(Deserialize, Clone)]
pub struct UpdateComplaintDTO {
    pub complaint_text: String,
}