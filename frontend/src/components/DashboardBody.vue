<script setup>
import { RouterLink } from 'vue-router';
</script>
<template>
  <!-- App body starts -->
  <div class="app-body">
    <div class="app-body">
      <!-- Row start -->
      <div class="col-xxl-12">
        <div class="card mb-3">
          <div class="card-body">
            <div class="table-responsive">
              <table class="table table-bordered m-0">
                <thead>
                  <tr>
                    <th>
                      <div class="d-flex justify-content-center">
                        <i class="bi bi-list fs-5"></i>
                      </div>
                    </th>
                    <th>Appraiser Name:</th>
                    <th>Member Name:</th>
                    <th>Dog Name:</th>
                    <th>Sex:</th>
                    <th>Age:</th>
                    <th>Weight:</th>
                    <th>Score:</th>
                    <th>MC#:</th>
                    <th>DNA#:</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td class="align-items-center">
                      <div class="d-flex justify-content-center">
                        <RouterLink to="/appraisal/asdfasd"><i class="bi bi-calculator fs-4"></i></RouterLink>
                      </div>
                    </td>
                    <td class="align-middle">Melinda</td>
                    <td class="align-middle">Michael Bogle</td>
                    <td class="align-middle">Getsuga</td>
                    <td class="align-middle">M</td>
                    <td class="align-middle">2yr</td>
                    <td class="align-middle">137lbs.</td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span class="badge border border-success text-success">87%</span>
                      </div>
                    </td>
                    <td class="align-middle">61685158</td>
                    <td class="align-middle">A19649A</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
      <!-- </div> -->
      <!-- Row end -->
    </div>
  </div>
  <!-- App body end -->
</template>
