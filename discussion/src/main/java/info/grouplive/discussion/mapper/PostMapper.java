package info.grouplive.discussion.mapper;

//import com.github.marlonlom.utilities.timeago.TimeAgo;
import info.grouplive.discussion.Repository.CommentRepository;
import info.grouplive.discussion.Repository.VoteRepository;
import info.grouplive.discussion.dto.PostRequest;
import info.grouplive.discussion.dto.PostResponse;
import info.grouplive.discussion.model.*;
import info.grouplive.discussion.service.AuthService;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.Optional;

import static info.grouplive.discussion.model.VoteType.DOWNVOTE;
import static info.grouplive.discussion.model.VoteType.UPVOTE;

@Mapper(componentModel = "spring")
public abstract class PostMapper {
    @Autowired
    private CommentRepository commentRepository;
    @Autowired
    private VoteRepository voteRepository;
    @Autowired
    private AuthService authService;

    @Mapping(target = "createdDate", expression = "java(java.time.Instant.now())")
    @Mapping(target = "description", source = "postRequest.description")
    @Mapping(target = "subreddit", source = "subreddit")
    @Mapping(target = "user", source = "user")
    @Mapping(target = "voteUpCount", constant = "0")
    @Mapping(target = "voteDownCount", constant = "0")
    public abstract Post map(PostRequest postRequest, Subreddit subreddit, UserModel user);

    @Mapping(target = "id", source = "post.postId")
    @Mapping(target = "postName", source = "post.postName")
    @Mapping(target = "description", source = "post.description")
    @Mapping(target = "url", source = "post.url")
    @Mapping(target = "subredditName", source = "post.subreddit.name")
    @Mapping(target = "userName", source = "post.user.username")
    @Mapping(target = "commentCount", expression = "java(commentCount(post))")
//    @Mapping(target = "duration", expression = "java(getDuration(post))")
    @Mapping(target = "upVote", expression = "java(isPostUpVoted(post, token))")
    @Mapping(target = "downVote", expression = "java(isPostDownVoted(post, token))")
    public abstract PostResponse mapToDto(Post post, String token);

    Integer commentCount(Post post) {
        return commentRepository.findByPost(post).size();
    }

//    String getDuration(Post post) {
//        return TimeAgo.using(post.getCreatedDate().toEpochMilli());
//    }

    boolean isPostUpVoted(Post post, String token) {
        return checkVoteType(post, UPVOTE, token);
    }

    boolean isPostDownVoted(Post post, String token) {
        return checkVoteType(post, DOWNVOTE, token);
    }

    private boolean checkVoteType(Post post, VoteType voteType, String token) {
        // TODO: replace with authService.getCurrentUser() after authService complete
        UserModel user = authService.getUser(token);
        Optional<Vote> voteForPostByUser =
                voteRepository.findTopByPostAndUserOrderByVoteIdDesc(post, user);
        return voteForPostByUser.filter(vote -> vote.getVoteType().equals(voteType))
                .isPresent();
    }
}