package com.grouplive.chat.repository;

import com.grouplive.chat.model.LoginUser;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface LoginUserRepository extends MongoRepository<LoginUser, String> {
    LoginUser findByUsername(String username);
}
