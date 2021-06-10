package info.grouplive.discussion.controller;

import info.grouplive.discussion.dto.PostRequest;
import info.grouplive.discussion.dto.PostResponse;
import info.grouplive.discussion.service.PostService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.util.List;


import static org.springframework.http.ResponseEntity.status;

@RestController
@RequestMapping("/api/posts/")
@AllArgsConstructor
public class PostController {

    private final PostService postService;

    @PostMapping
    public ResponseEntity<Void> createPost(@RequestHeader(value="Authorization") String token, @RequestBody PostRequest postRequest) {
        postService.save(postRequest, token);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @GetMapping
    public ResponseEntity<List<PostResponse>> getAllPosts(@RequestHeader(value="Authorization") String token) {
        return status(HttpStatus.OK).body(postService.getAllPosts(token));
    }

    @GetMapping("/{id}")
    public ResponseEntity<PostResponse> getPost(@RequestHeader(value="Authorization") String token, @PathVariable Long id) {
        return status(HttpStatus.OK).body(postService.getPost(id, token));
    }

    @GetMapping("by-subreddit/{id}")
    public ResponseEntity<List<PostResponse>> getPostsBySubreddit(@RequestHeader(value="Authorization") String token, Long id) {
        return status(HttpStatus.OK).body(postService.getPostsBySubreddit(id, token));
    }

    @GetMapping("by-user/{name}")
    public ResponseEntity<List<PostResponse>> getPostsByUsername(@RequestHeader(value="Authorization") String token, @PathVariable String name) {
        return status(HttpStatus.OK).body(postService.getPostsByUsername(name, token));
    }
}