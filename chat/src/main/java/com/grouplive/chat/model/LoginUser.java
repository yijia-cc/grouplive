package com.grouplive.chat.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.mongodb.core.mapping.Document;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Document
public class LoginUser {

    private String id;
    private String username;
    private String lastname;
    private String firstname;
    private String email;
    private String phone;
}