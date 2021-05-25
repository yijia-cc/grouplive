package info.grouplive.discussion.Repository;


import info.grouplive.discussion.model.Comment;
import info.grouplive.discussion.model.Post;
import info.grouplive.discussion.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CommentRepository extends JpaRepository<Comment, Long> {
    List<Comment> findByPost(Post post);

    List<Comment> findAllByUser(User user);
}
