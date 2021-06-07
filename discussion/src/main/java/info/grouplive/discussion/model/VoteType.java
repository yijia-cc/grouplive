package info.grouplive.discussion.model;

import info.grouplive.discussion.exceptions.DiscussionException;
import java.util.Arrays;

public enum VoteType {
    UPVOTE(1), DOWNVOTE(-1), NULLVOTE(0);

    private int direction;

    VoteType(int direction) {
    }

    public static VoteType lookup(Integer direction) {
        return Arrays.stream(VoteType.values())
                .filter(value -> value.getDirection().equals(direction))
                .findAny()
                .orElseThrow(() -> new DiscussionException("Vote not found"));
    }

    public Integer getDirection() {
        return direction;
    }
}
