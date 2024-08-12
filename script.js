import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '50s', target: 50000 }, // 50,000 usuarios virtuales en 30 segundos
  ],
};

export default function () {
  let res = http.get('http://localhost:8080/categories/MLA5725');
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1);
}
