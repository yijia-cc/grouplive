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
import info.grouplive.discussion.model.User;
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

    public Post save(PostRequest postRequest) {
        Subreddit subreddit = subredditRepository.findByName(postRequest.getSubredditName())
                .orElseThrow(() -> new SubredditNotFoundException(postRequest.getSubredditName()));
        // TODO: replace with authService.getCurrentUser() after authService complete
        User currentUser = new User(1l, "admin", "123", "admin@gmail.com", null, true);
        postRepository.save(postMapper.map(postRequest, subreddit, currentUser));
        return postMapper.map(postRequest, subreddit, currentUser);
    }

    @Transactional(readOnly = true)
    public PostResponse getPost(Long id) {
        Post post = postRepository.findById(id)
                .orElseThrow(() -> new PostNotFoundException(id.toString()));
        return postMapper.mapToDto(post);
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getAllPosts() {
         return postRepository.findAll(Sort.by(Sort.Direction.DESC, "postId"))
                                    .stream()
                                    .map(postMapper::mapToDto)
                                    .collect(toList());
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getPostsBySubreddit(Long subredditId) {
        Subreddit subreddit = subredditRepository.findById(subredditId)
                .orElseThrow(() -> new SubredditNotFoundException(subredditId.toString()));
        List<Post> posts = postRepository.findAllBySubreddit(subreddit);
        return posts.stream().map(postMapper::mapToDto).collect(toList());
    }

    @Transactional(readOnly = true)
    public List<PostResponse> getPostsByUsername(String username) {
        User user = userRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException(username));
        return postRepository.findByUser(user)
                .stream()
                .map(postMapper::mapToDto)
                .collect(toList());
    }
}
