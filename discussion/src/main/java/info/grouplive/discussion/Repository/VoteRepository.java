package info.grouplive.discussion.Repository;

import info.grouplive.discussion.model.Post;
import info.grouplive.discussion.model.UserModel;
import info.grouplive.discussion.model.Vote;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface VoteRepository extends JpaRepository<Vote, Long> {
    Optional<Vote> findTopByPostAndUserOrderByVoteIdDesc(Post post, UserModel currentUser);
}