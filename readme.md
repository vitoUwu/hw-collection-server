# HW Collection API Documentation

## Overview

This API provides endpoints to retrieve and search for cars, allowing users to get details about all available cars or to search for specific models.

## Endpoints

### 1. Get All Cars

- **GET** `/cars`
- **Summary**: Retrieve a list of all cars
- **Responses**:
    - **200 OK**:
      ```json
      [
          {
              "base_color": "string",
              "code": "string",
              "col_code": "string",
              "color": "string",
              "country": "string",
              "created_at": "string",
              "id": "string",
              "image": "string",
              "interior_color": "string",
              "model_name": "string",
              "notes": "string",
              "series_code": "string",
              "series_name": "string",
              "tampo": "string",
              "updated_at": "string",
              "wheel_type": "string",
              "window_color": "string"
          }
      ]
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

### 2. Search Cars

- **GET** `/search`
- **Summary**: Search for cars based on a query
- **Query Parameters**:
    - `q` (required): The search query (string)
- **Responses**:
    - **200 OK**:
      ```json
      [
          {
              "code": "string",
              "col_code": "string",
              "id": "string",
              "image": "string",
              "model_name": "string",
              "notes": "string",
              "series_code": "string",
              "series_name": "string"
          }
      ]
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

## Error Responses

All error responses will follow the format:
```json
{
    "error": "string"
}
