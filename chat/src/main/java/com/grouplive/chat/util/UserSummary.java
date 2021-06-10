package com.grouplive.chat.util;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class UserSummary{
    private String id;
    private String name;
    private String profilePicture;
}