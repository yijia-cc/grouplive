package com.amr.chatservice.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import javax.persistence.Id;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Table;
import java.util.Date;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
//@Table(name = "chatMessage")
public class ChatMessage {
   @Id
   @GeneratedValue(strategy = GenerationType.AUTO)
   private String id;
   private String chatId;
   private String senderId;
   private String recipientId;
   private String senderName;
   private String recipientName;
   private String content;
   private Date timestamp;
   private MessageStatus status;
}
