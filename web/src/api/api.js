var APIURL = "http://localhost:8448";

export function callAPI(url, options) {
  return new Promise((resolve, reject) => {
    fetch(url, options)
      .then((res) => {
        // successful api call
        if (res.status >= 200 && res.status < 300) {
          resolve(res);
          return;
        }

        // 400 are usually a bad request, let client know to display message
        if (res.status === 400) {
          resolve(res);
          return;
        }
        reject(res);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export function callJSONApi(url, options) {
  return new Promise((resolve, reject) => {
    callAPI(url, options)
      .then((res) => parseJSON(res))
      .then((data) => {
        if (data?.error) {
          reject(data);
        }
        resolve(data);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export function parseJSON(response) {
  return response.text().then((text) => {
    if (!text) {
      return {};
    }
    return JSON.parse(text);
  });
}

export function buildOptions(method, payload) {
  const options = {
    method: method.toUpperCase(),
    headers: {
      Accept: "application/json",
    },
    mode: "cors",
  };

  if (method.toLowerCase() !== "get") {
    options.headers["Content-Type"] = "application/json";
  }

  if (payload) {
    options.body = JSON.stringify(payload);
  }

  return options;
}

export function postLogin(user) {
  const url = `${APIURL}/login`;
  const options = buildOptions("POST", user);

  return callJSONApi(url, options);
}

export function postRegistration(user) {
  const url = `${APIURL}/register`;
  const options = buildOptions("POST", user);

  return callJSONApi(url, options);
}

export function getUser(userId) {
  const url = `${APIURL}/user/${userId}`;
  const options = buildOptions("GET");

  return callJSONApi(url, options);
}
