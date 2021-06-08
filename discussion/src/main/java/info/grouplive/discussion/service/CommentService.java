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

    public void save(CommentsDto commentsDto) {
        Post post = postRepository.findById(commentsDto.getPostId())
                .orElseThrow(() -> new PostNotFoundException(commentsDto.getPostId().toString()));
        // TODO: replace with authService.getCurrentUser() after authService complete
        User currentUser = new User(1l, "admin", "123", "admin@gmail.com", null, true);
        Comment comment = commentMapper.map(commentsDto, post, currentUser);
        commentRepository.save(comment);
    }

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
