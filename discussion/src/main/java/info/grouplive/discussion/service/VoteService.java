package info.grouplive.discussion.service;

import info.grouplive.discussion.Repository.PostRepository;
import info.grouplive.discussion.Repository.VoteRepository;
import info.grouplive.discussion.dto.VoteDto;
import info.grouplive.discussion.exceptions.PostNotFoundException;
import info.grouplive.discussion.model.Post;
import info.grouplive.discussion.model.User;
import info.grouplive.discussion.model.Vote;
import lombok.AllArgsConstructor;
import org.springframework.security.authentication.DisabledException;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

import static info.grouplive.discussion.model.VoteType.UPVOTE;

@Service
@AllArgsConstructor
public class VoteService {
    private final VoteRepository voteRepository;
    private final PostRepository postRepository;

    @Transactional
    public void vote(VoteDto voteDto) {
        Post post = postRepository.findById(voteDto.getPostId())
                .orElseThrow(() -> new PostNotFoundException("Post Not Found with ID - " + voteDto.getPostId()));
        // TODO: replace with authService.getCurrentUser() after authService complete
        User currentUser = new User(1l, "admin", "123", "admin@gmail.com", null, true);
        Optional<Vote> voteByPostAndUser = voteRepository.findTopByPostAndUserOrderByVoteIdDesc(post, currentUser);
        if (voteByPostAndUser.isPresent() &&
                voteByPostAndUser.get().getVoteType()
                        .equals(voteDto.getVoteType())) {
            throw new DisabledException("You have already "
                    + voteDto.getVoteType() + "'d for this post");
        }
        if (UPVOTE.equals(voteDto.getVoteType())) {
            post.setVoteCount(post.getVoteCount() + 1);
        } else {
            post.setVoteCount(post.getVoteCount() - 1);
        }
        voteRepository.save(mapToVote(voteDto, post));
        postRepository.save(post);
    }

    private Vote mapToVote(VoteDto voteDto, Post post) {
        // TODO: replace with authService.getCurrentUser() after authService complete
        User currentUser = new User(1l, "admin", "123", "admin@gmail.com", null, true);
        return Vote.builder()
                .voteType(voteDto.getVoteType())
                .post(post)
                .user(currentUser)
                .build();
    }
}
