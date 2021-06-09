package info.grouplive.discussion.controller;

import info.grouplive.discussion.dto.PostResponse;
import info.grouplive.discussion.dto.VoteDto;
import info.grouplive.discussion.service.VoteService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/votes/")
@AllArgsConstructor
public class VoteController {
    private final VoteService voteService;

    @PostMapping
    public ResponseEntity<List<PostResponse>> vote(@RequestHeader(value="Authorization") String token, @RequestBody VoteDto voteDto) {
        voteService.vote(voteDto, token);
        return new ResponseEntity<>(HttpStatus.OK);
    }
}
