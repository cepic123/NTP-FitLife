use actix_web::{get, post, put, delete, web, Responder, HttpResponse};
use crate::{AppState, Complaint};
use super::models::{CreateComplaintDTO, UpdateComplaintDTO};

#[get("/complaint")]
async fn get_complaints(data: web::Data<AppState>) -> impl Responder {
    HttpResponse::Ok().json(data.complaints.lock().unwrap().to_vec())
}

#[post("/complaint")]
async fn create_complaint(data: web::Data<AppState>, body: web::Json<CreateComplaintDTO>) -> impl Responder {
    let mut complaints = data.complaints.lock().unwrap();
    let new_id = complaints.len() + 1; 
    complaints.push(Complaint {
        id: new_id,
        user_id: body.user_id,
        complaint_subject_id: body.complaint_subject_id,
        complaint_text: body.complaint_text.clone(),
    });

    HttpResponse::Ok().json(complaints.to_vec())
}

#[put("/complaint{id}")]
async fn update_complaint(data: web::Data<AppState>, path: web::Path<i32>, param_obj: web::Json<UpdateComplaintDTO>) -> impl Responder {
    let id = path.into_inner();
    let complaints = data.complaints.lock().unwrap();

    HttpResponse::Ok().json(complaints.to_vec())
}


#[delete("/complaint{id}")]
async fn delete_complaint(data: web::Data<AppState>, path: web::Path<i32>) -> impl Responder {
    let id = path.into_inner();

    HttpResponse::Ok().json(id)
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(get_complaints)
    .service(create_complaint)
    .service(update_complaint)
    .service(delete_complaint);
}
