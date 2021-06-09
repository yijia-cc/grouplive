package info.grouplive.discussion.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotEmpty;
import java.time.Instant;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class UserModel {
    @Id
    @SequenceGenerator(name = "userId", sequenceName = "octo_reference_code", allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "octo_reference_code")
    private String userId;
    @NotBlank(message = "Username is required")
    private String username;
    private String lastName;
    private String firstName;
    @Email
    @NotEmpty(message = "Email is required")
    private String email;
    private String phone;
    private String address;
    private String aptNumber;
}

