import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "10s", target: 25 },
    { duration: "20s", target: 50 },
    { duration: "20s", target: 100 },
    { duration: "10s", target: 0 },
  ],

  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<500"],
  },
};

const BASE_URL = "http://localhost:8080";

export default function () {
  const loginRes = http.post(
    `${BASE_URL}/login`,
    JSON.stringify({
      email: "benchmark@test.com",
      password: "benchmark123",
    }),
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  check(loginRes, {
    "login successful": (r) => r.status === 200,
  });

  if (loginRes.status !== 200) {
    return;
  }

  const token = loginRes.json("token");

  const notesRes = http.get(
    `${BASE_URL}/notes?page=1&limit=10`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  check(notesRes, {
    "get notes successful": (r) => r.status === 200,
  });

  sleep(1);
}
