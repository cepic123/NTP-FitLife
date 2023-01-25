use actix_web::{get, post, put, delete, web::{Data, Json, Path, ServiceConfig}, Responder, HttpResponse};
use crate::{AppState, Complaint};
use super::models::{CreateComplaintDTO, UpdateComplaintDTO};
use sqlx::{self, FromRow};
use serde::{Deserialize, Serialize};


#[put("/complaint/{id}")]
async fn update_complaint(state: Data<AppState>, path: Path<i32>, body: Json<UpdateComplaintDTO>) -> impl Responder {

    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Complaint>(
        "UPDATE complaints SET complaint_text = $1 WHERE id = $2"
    )
        .bind(body.complaint_text.to_string())
        .bind(id)
        .fetch_one(&state.db)
        .await
    {
        Ok(complaint) => HttpResponse::Ok().json(complaint),
        Err(_) => HttpResponse::NotFound().json("Complaint update not succesfull"),
    }
}

#[get("/complaint/{id}")]
async fn get_complaint(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Complaint>(
        "SELECT id, user_id, user_name, subject_name, complaint_subject_id, complaint_text FROM complaints WHERE id = $1"
    )
        .bind(id)
        .fetch_one(&state.db)
        .await
    {
        Ok(complaints) => HttpResponse::Ok().json(complaints),
        Err(_) => HttpResponse::NotFound().json("No complaint found"),
    }
}

#[get("/complaint")]
async fn get_complaints(state: Data<AppState>) -> impl Responder {
    println!("hello there!");

    match sqlx::query_as::<_,Complaint>(
        "SELECT id, user_id, user_name, subject_name, complaint_subject_id, complaint_text FROM complaints"
    )
        .fetch_all(&state.db)
        .await
    {
        Ok(complaints) => HttpResponse::Ok().json(complaints),
        Err(_) => HttpResponse::NotFound().json("No complaints found"),
    }
}

#[get("/complaint/user/{id}")]
async fn get_user_complaints(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Complaint>(
        "SELECT id, user_id, user_name, subject_name, complaint_subject_id, complaint_text FROM complaints WHERE user_id = $1"
    )
        .bind(id)
        .fetch_all(&state.db)
        .await
    {
        Ok(complaints) => HttpResponse::Ok().json(complaints),
        Err(_) => HttpResponse::NotFound().json("No complaints found by user"),
    }
}

#[get("/complaint/subject/{id}")]
async fn get_subject_complaints(state: Data<AppState>, path: Path<i32>) -> impl Responder {
    
    let id: i32 = path.into_inner();

    match sqlx::query_as::<_,Complaint>(
        "SELECT id, user_id, user_name, subject_name, complaint_subject_id, complaint_text FROM complaints WHERE complaint_subject_id = $1"
    )
        .bind(id)
        .fetch_all(&state.db)
        .await
    {
        Ok(complaints) => HttpResponse::Ok().json(complaints),
        Err(_) => HttpResponse::NotFound().json("No complaints found for subject"),
    }
}

#[post("/complaint")]
async fn create_complaint(state: Data<AppState>, body: Json<CreateComplaintDTO>) -> impl Responder {
    println!("hello there!");
    match sqlx::query_as::<_, Complaint>(
        "INSERT INTO complaints (user_id, user_name, subject_name, complaint_subject_id, complaint_text) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, user_name, subject_name, complaint_subject_id, complaint_text"
    )
        .bind(body.user_id)
        .bind(body.user_name.to_string())
        .bind(body.subject_name.to_string())
        .bind(body.complaint_subject_id)
        .bind(body.complaint_text.to_string())
        .fetch_one(&state.db)
        .await
        {
            Ok(complaint) => HttpResponse::Ok().json(complaint),
            Err(_) => HttpResponse::InternalServerError().json("Failed to create user"),
        }
}

pub fn config(cfg: &mut ServiceConfig) {
    cfg.service(get_complaint)
    .service(update_complaint)
    .service(get_complaints)
    .service(create_complaint)
    .service(get_subject_complaints)
    .service(get_user_complaints);
}
