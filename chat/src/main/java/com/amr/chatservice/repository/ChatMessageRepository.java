package com.amr.chatservice.repository;

import com.amr.chatservice.model.ChatMessage;
import com.amr.chatservice.model.MessageStatus;
//import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.CrudRepository;


import java.util.List;

public interface ChatMessageRepository
        extends CrudRepository<ChatMessage, String> {

    long countBySenderIdAndRecipientIdAndStatus(
            String senderId, String recipientId, MessageStatus status);

    List<ChatMessage> findByChatId(String chatId);
    ChatMessage findBySenderIdAndRecipientId(String senderId, String recipientId);
}