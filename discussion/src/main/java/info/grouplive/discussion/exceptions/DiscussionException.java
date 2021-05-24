package info.grouplive.discussion.exceptions;

public class DiscussionException extends RuntimeException {
    public DiscussionException(String exMessage, Exception exception) {
        super(exMessage, exception);
    }

    public DiscussionException(String exMessage) {
        super(exMessage);
    }
}
