package info.grouplive.discussion.service;


import info.grouplive.discussion.Repository.CommentRepository;
import info.grouplive.discussion.Repository.PostRepository;
import info.grouplive.discussion.Repository.UserRepository;
import info.grouplive.discussion.dto.CommentsDto;
import info.grouplive.discussion.exceptions.PostNotFoundException;
import info.grouplive.discussion.mapper.CommentMapper;
import info.grouplive.discussion.model.Comment;
import info.grouplive.discussion.model.Post;
import info.grouplive.discussion.model.User;
import lombok.AllArgsConstructor;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.List;

import static java.util.stream.Collectors.toList;

@Service
@AllArgsConstructor
public class CommentService {
    private static final String POST_URL = "";
    private final PostRepository postRepository;
    private final UserRepository userRepository;
    private final CommentMapper commentMapper;
    private final CommentRepository commentRepository;
//    private final AuthService authService;
//    private final MailContentBuilder mailContentBuilder;
//    private final MailService mailService;

    public void save(CommentsDto commentsDto) {
        Post post = postRepository.findById(commentsDto.getPostId())
                .orElseThrow(() -> new PostNotFoundException(commentsDto.getPostId().toString()));
        Comment comment = commentMapper.map(commentsDto, post, new User(123l, "admin", "123", "", null, true)); // authService.getCurrentUser()
        commentRepository.save(comment);

//        String message = mailContentBuilder.build(authService.getCurrentUser() + " posted a comment on your post." + POST_URL);
//        sendCommentNotification(message, post.getUser());
    }

//    private void sendCommentNotification(String message, User user) {
//        mailService.sendMail(new NotificationEmail(user.getUsername() + " Commented on your post", user.getEmail(), message));
//    }

    public List<CommentsDto> getAllCommentsForPost(Long postId) {
        Post post = postRepository.findById(postId).orElseThrow(() -> new PostNotFoundException(postId.toString()));
        return commentRepository.findByPost(post)
                    .stream()
                    .map(commentMapper::mapToDto).collect(toList());
    }

    public List<CommentsDto> getAllCommentsForUser(String userName) {
        User user = userRepository.findByUsername(userName)
                            .orElseThrow(() -> new UsernameNotFoundException(userName));
        return commentRepository.findAllByUser(user)
                .stream()
                .map(commentMapper::mapToDto)
                .collect(toList());
    }
}
