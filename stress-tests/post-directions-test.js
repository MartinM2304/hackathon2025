import http from "k6/http";
import { sleep } from "k6";

export const options = {
  iterations: 30,
};

export default function () {
  http.post("http://localhost:3000/api/direction", {
    direction: 2,
  });

  sleep(0.1);
}
