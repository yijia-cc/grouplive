package com.grouplive.chat.service;

import com.grouplive.chat.model.LoginUser;
import com.grouplive.chat.repository.LoginUserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {
    @Autowired
    private LoginUserRepository loginUserRepository;
    public List<LoginUser> findAll(){
        return loginUserRepository.findAll();
    }

}
