package info.grouplive.discussion.service;

import info.grouplive.discussion.Repository.PostRepository;
import info.grouplive.discussion.Repository.SubredditRepository;
import info.grouplive.discussion.Repository.UserRepository;
import info.grouplive.discussion.dto.PostRequest;
import info.grouplive.discussion.dto.PostResponse;
import info.grouplive.discussion.exceptions.PostNotFoundException;
import info.grouplive.discussion.exceptions.SubredditNotFoundException;
import info.grouplive.discussion.mapper.PostMapper;
import info.grouplive.discussion.model.Post;
import info.grouplive.discussion.model.Subreddit;
//import info.grouplive.discussion.model.User;
import info.grouplive.discussion.model.UserModel;
import io.grpc.StatusRuntimeException;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

import static java.util.stream.Collectors.toList;

@Service
@AllArgsConstructor
@Slf4j
@Transactional
public class PostService {

    private final PostRepository postRepository;
    private final SubredditRepository subredditRepository;
    private final UserRepository userRepository;
    private final PostMapper postMapper;
    private final AuthService authService;

    public Post save(PostRequest postRequest, String token) {
        Subreddit subreddit = subredditRepository.findByName(postRequest.getSubredditName())
                .orElseThrow(() -> new SubredditNotFoundException(postRequest.getSubredditName()));
        try {
            UserModel user = authService.getUser(token);
            postRepository.save(postMapper.map(postRequest, subreddit, user));
            return postMapper.map(postRequest, subreddit, user);
        } catch (StatusRuntimeException e) {
            e.printStackTrace();
        }
        return null;
    }

    @Transactional(readOnly = true)
    public PostResponse getPost(Long id, String token) {
        Post post = postRepository.findById(id)
                .orElseThrow(() -> new PostNotFoundException(id.toString()));
        return postMapper.mapToDto(post, token);
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getAllPosts(String token) {
         return postRepository.findAll(Sort.by(Sort.Direction.DESC, "postId"))
                                    .stream()
                                    .map((post) -> postMapper.mapToDto(post, token))
                                    .collect(toList());
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getPostsBySubreddit(Long subredditId, String token) {
        Subreddit subreddit = subredditRepository.findById(subredditId)
                .orElseThrow(() -> new SubredditNotFoundException(subredditId.toString()));
        List<Post> posts = postRepository.findAllBySubreddit(subreddit);
        return posts.stream().map((post) -> postMapper.mapToDto(post, token)).collect(toList());
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getPostsByUsername(String username, String token) {
        UserModel user = userRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException(username));
        return postRepository.findByUser(user)
                .stream()
                .map((post) -> postMapper.mapToDto(post, token))
                .collect(toList());
    }
}
