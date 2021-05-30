package com.grouplive.chat;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@SpringBootApplication
public class ChatApplication {
	public static void main(String[] args) {
		SpringApplication.run(ChatApplication.class, args);
	}

	@GetMapping("/chat")
	public String hello() {
	    return "Welcome to chat!";
	}
}
