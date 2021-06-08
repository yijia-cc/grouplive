package info.grouplive.discussion.controller;

import info.grouplive.discussion.dto.CommentsDto;
import info.grouplive.discussion.model.Comment;
import info.grouplive.discussion.service.CommentService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

import static org.springframework.http.HttpStatus.CREATED;
import static org.springframework.http.HttpStatus.OK;
import static org.springframework.http.ResponseEntity.status;

@RestController
@RequestMapping("/api/comments/")
@AllArgsConstructor
public class CommentsController {
    private final CommentService commentService;

    @PostMapping
    public ResponseEntity<List<CommentsDto>> createComment(@RequestBody CommentsDto commentsDto) {
        commentService.save(commentsDto);
        return getAllCommentsForPost(commentsDto.getPostId());
    }

    @GetMapping("/by-post/{postId}")
    public ResponseEntity<List<CommentsDto>> getAllCommentsForPost(@PathVariable Long postId) {
        return ResponseEntity.status(OK)
                    .body(commentService.getAllCommentsForPost(postId));
    }

    @GetMapping("/by-user/{userName}")
    public ResponseEntity<List<CommentsDto>> getAllCommentsForUser(@PathVariable String userName) {
        return ResponseEntity.status(OK)
                    .body(commentService.getAllCommentsForUser(userName));
    }
}
