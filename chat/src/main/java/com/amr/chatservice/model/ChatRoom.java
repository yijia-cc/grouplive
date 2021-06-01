package com.amr.chatservice.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;


@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
//@Table(name = "chatRoom")
public class ChatRoom {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private String id;
    private String chatId;
    private String senderId;
    private String recipientId;
}
