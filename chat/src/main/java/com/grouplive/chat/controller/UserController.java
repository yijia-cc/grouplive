package com.grouplive.chat.controller;

import com.grouplive.chat.model.LoginUser;
import com.grouplive.chat.repository.LoginUserRepository;
import com.grouplive.chat.service.UserService;
import com.grouplive.chat.util.UserSummary;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;

@RestController
public class UserController {
    @Autowired
    private LoginUserRepository loginUserRepository;
    private UserService userService;
    @GetMapping(value = "/users/me", produces = MediaType.APPLICATION_JSON_VALUE)
    public UserSummary getCurrentUser() {
        LoginUser loginUser = LoginUser
                .builder()
                .id("u08")
                .username("yuran300")
                .lastname("Liu2")
                .firstname("yuran2")
                .email("1232")
                .phone("121321343")
                .build();

        loginUserRepository.save(loginUser);
        LoginUser loginUser2 = loginUserRepository.findByUsername("yuran300");

//        LoginUser loginUser = LoginUser
//                .builder()
//                .id("u09")
//                .username("yuran200")
//                .lastname("Liu")
//                .firstname("yuran")
//                .email("12321")
//                .phone("12132132132")
//                .build();

        return UserSummary
                .builder()
                .id(loginUser2.getId())
                .name(loginUser2.getUsername())
                .profilePicture("https://www.w3schools.com/w3images/avatar5.png")
                .build();
    }

    @GetMapping(value = "/users/summaries", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<?> findAllUserSummaries(){
        List<LoginUser> users = new ArrayList<>();
        LoginUser u1 = LoginUser
                .builder()
                .id("u09")
                .username("yuran200")
                .lastname("Liu")
                .firstname("yuran")
                .email("12321")
                .phone("12132132132")
                .build();
        LoginUser u2 = LoginUser
                .builder()
                .id("u07")
                .username("yuran400")
                .lastname("Liu3")
                .firstname("yuran3")
                .email("1232")
                .phone("12132132")
                .build();
//        users.add(u1);
//        users.add(u2);
//
        loginUserRepository.save(u1);
        loginUserRepository.save(u2);

        //return ResponseEntity.ok(users.stream().map(this::convertTo));
        return ResponseEntity.ok(loginUserRepository
                .findAll()
                .stream()
                .filter(user -> !user.getUsername().equals(getCurrentUser().getName()))
                .map(this::convertTo));
    }

    private UserSummary convertTo(LoginUser user) {
        return UserSummary
                .builder()
                .id(user.getId())
                .name(user.getUsername())
                .profilePicture("https://www.w3schools.com/howto/img_avatar.png")
                .build();
    }

}
