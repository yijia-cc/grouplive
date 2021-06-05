const AUTH_SERVICE = "http://localhost:8081";
const CHAT_SERVICE = "http://localhost:8080";

const token = "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dXJhbjIiLCJhdXRob3JpdGllcyI6WyJST0xFX1VTRVIiXSwiaWF0IjoxNjIyODMzMDY1LCJleHAiOjE2MjI5MTk0NjV9.-RYv_DaPX2tACz2O_0erzhE-fYGmSGk94aj_YGpQQKd-JXd5HZ-V5rR_GVwewu2YwF9Dr7WPS3rv7FX0QmUJIg";

const request = (options) => {
  const headers = new Headers();

  if (options.setContentType !== false) {
    headers.append("Content-Type", "application/json");
  }

  if (token) {
    headers.append(
      "Authorization",
      "Bearer " + token
    );
  }

  const defaults = { headers: headers };
  options = Object.assign({}, defaults, options);

  return fetch(options.url, options).then((response) =>
    response.json().then((json) => {
      if (!response.ok) {
        return Promise.reject(json);
      }
      return json;
    })
  );
};

// export function login(loginRequest) {
//   return request({
//     url: AUTH_SERVICE + "/signin",
//     method: "POST",
//     body: JSON.stringify(loginRequest),
//   });
// }

// export function facebookLogin(facebookLoginRequest) {
//   return request({
//     url: AUTH_SERVICE + "/facebook/signin",
//     method: "POST",
//     body: JSON.stringify(facebookLoginRequest),
//   });
// }

// export function signup(signupRequest) {
//   return request({
//     url: AUTH_SERVICE + "/users",
//     method: "POST",
//     body: JSON.stringify(signupRequest),
//   });
// }

export function getCurrentUser() {
  if (!token) {
    return Promise.reject("No access token set.");
  }

  return new Promise((resolve, reject) => {
    resolve(
      {
      "id": "60ba77a259159331b873c1c4", 
      "username": "yuran2", 
      "name": "yuran2", 
      "profilePicture": "https://www.w3schools.com/w3images/avatar5.png"}
    )
  });
}

export function getUsers() {
  if (!token) {
    return Promise.reject("No access token set.");
  }
  return new Promise((resolve, reject) => {
    resolve([
      {
      "id": "60b8072e5915933c8cd8df45", 
      "username": "yuran300", 
      "name": "yuran1300", 
      "profilePicture": "https://www.w3schools.com/howto/img_avatar2.png"
      },
      {
      "id": "60b8072e5915933c8cd8df46", 
      "username": "yuran400", 
      "name": "yuran1400", 
      "profilePicture": "https://www.w3schools.com/howto/img_avatar.png"
      },
      {
        "id": "60b8072e5915933c8cd8df47", 
        "username": "yuran500", 
        "name": "yuran1500", 
        "profilePicture": "https://www.w3schools.com/w3images/avatar6.png"
        }
      ])
  });
}

export function countNewMessages(senderId, recipientId) {
  if (!token) {
    return Promise.reject("No access token set.");
  }

  return request({
    url: CHAT_SERVICE + "/messages/" + senderId + "/" + recipientId + "/count",
    method: "GET",
  });
}

export function findChatMessages(senderId, recipientId) {
  if (!token) {
    return Promise.reject("No access token set.");
  }

  return request({
    url: CHAT_SERVICE + "/messages/" + senderId + "/" + recipientId,
    method: "GET",
  });
}

export function findChatMessage(id) {
  if (!token) {
    return Promise.reject("No access token set.");
  }

  return request({
    url: CHAT_SERVICE + "/messages/" + id,
    method: "GET",
  });
}
