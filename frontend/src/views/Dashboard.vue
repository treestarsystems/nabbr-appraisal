<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { SwalToastError } from '../helpers/sweetalert';
import { loadThirdPartyJS } from '../helpers/script';
import { thirdPartyJSFilePaths } from '../helpers/thirdPartyFIlesList';

onMounted(async () => {
  const authStore = useAuthStore();
  const swal: any = inject('$swal');
  // Check for non-null user
  await authStore.protectView;
  // Check for unprivileged user
  const authorizedPrivilegeLevels: string[] = ['ADMIN', 'APPRAISER'];
  const isAuthorized = authStore.checkUserPrivilegeLevel(authorizedPrivilegeLevels);
  if (!isAuthorized) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastError(swal, 'User Unauthorized ');
  }
  loadThirdPartyJS(thirdPartyJSFilePaths);
});
</script>

<template>
  <!-- Page wrapper start -->
  <div class="page-wrapper">
    <!-- Main container start -->
    <div class="main-container">
      <!-- Sidebar wrapper start -->
      <nav id="sidebar" class="sidebar-wrapper">
        <!-- App brand starts -->
        <div class="app-brand px-3 py-2 d-flex align-items-center">
          <a href="index.html">
            <img src="/dog.svg" class="logo" alt="Bootstrap Gallery" />
          </a>
        </div>
        <!-- App brand ends -->

        <!-- Sidebar profile starts -->
        <div class="sidebar-profile">
          <img src="/assets/images/user1.png" class="img-3x me-3 rounded-3" alt="Admin Dashboard" />
          <div class="m-0">
            <p class="m-0">Hello &#128075;</p>
            <h6 class="m-0 text-nowrap">Mr. Yuki Michiels</h6>
          </div>
        </div>
        <!-- Sidebar profile ends -->

        <!-- Sidebar menu starts -->
        <div class="sidebarMenuScroll">
          <ul class="sidebar-menu">
            <li class="active current-page">
              <a href="index.html">
                <i class="bi bi-pie-chart"></i>
                <span class="menu-text">Dashboard</span>
              </a>
            </li>
            <li>
              <a href="analytics.html">
                <i class="bi bi-bar-chart-line"></i>
                <span class="menu-text">Analytics</span>
              </a>
            </li>
            <li>
              <a href="widgets.html">
                <i class="bi bi-box"></i>
                <span class="menu-text">Widgets</span>
              </a>
            </li>
            <li>
              <a href="calendar.html">
                <i class="bi bi-calendar2"></i>
                <span class="menu-text">Calendar</span>
              </a>
            </li>
            <li>
              <a href="default.html">
                <i class="bi bi-layout-sidebar"></i>
                <span class="menu-text">Default Layout</span>
              </a>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-stickies"></i>
                <span class="menu-text">Components</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="accordions.html">Accordions</a>
                </li>
                <li>
                  <a href="alerts.html">Alerts</a>
                </li>
                <li>
                  <a href="buttons.html">Buttons</a>
                </li>
                <li>
                  <a href="badges.html">Badges</a>
                </li>
                <li>
                  <a href="carousel.html">Carousel</a>
                </li>
                <li>
                  <a href="list-items.html">List Items</a>
                </li>
                <li>
                  <a href="progress.html">Progress Bars</a>
                </li>
                <li>
                  <a href="popovers.html">Popovers</a>
                </li>
                <li>
                  <a href="tooltips.html">Tooltips</a>
                </li>
              </ul>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-ui-checks-grid"></i>
                <span class="menu-text">Forms</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="form-inputs.html">Form Inputs</a>
                </li>
                <li>
                  <a href="form-checkbox-radio.html">Checkbox &amp; Radio</a>
                </li>
                <li>
                  <a href="form-file-input.html">File Input</a>
                </li>
                <li>
                  <a href="form-validations.html">Validations</a>
                </li>
                <li>
                  <a href="date-time-pickers.html">Date Time Pickers</a>
                </li>
                <li>
                  <a href="form-layouts.html">Form Layouts</a>
                </li>
              </ul>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-window-sidebar"></i>
                <span class="menu-text">Invoices</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="create-invoice.html">Create Invoice</a>
                </li>
                <li>
                  <a href="view-invoice.html">View Invoice</a>
                </li>
                <li>
                  <a href="invoice-list.html">Invoice List</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="tables.html">
                <i class="bi bi-border-all"></i>
                <span class="menu-text">Tables</span>
              </a>
            </li>
            <li>
              <a href="events.html">
                <i class="bi bi-calendar4"></i>
                <span class="menu-text">Events</span>
              </a>
            </li>
            <li>
              <a href="subscribers.html">
                <i class="bi bi-check-circle"></i>
                <span class="menu-text">Subscribers</span>
              </a>
            </li>
            <li>
              <a href="contacts.html">
                <i class="bi bi-wallet2"></i>
                <span class="menu-text">Contacts</span>
              </a>
            </li>
            <li>
              <a href="settings.html">
                <i class="bi bi-gear"></i>
                <span class="menu-text">Settings</span>
              </a>
            </li>
            <li>
              <a href="profile.html">
                <i class="bi bi-person-square"></i>
                <span class="menu-text">Profile</span>
              </a>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-code-square"></i>
                <span class="menu-text">Cards</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="cards.html">Cards</a>
                </li>
                <li>
                  <a href="advanced-cards.html">Advanced Cards</a>
                </li>
              </ul>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-pie-chart"></i>
                <span class="menu-text">Graphs</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="apex.html">Apex</a>
                </li>
                <li>
                  <a href="morris.html">Morris</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="maps.html">
                <i class="bi bi-pin-map"></i>
                <span class="menu-text">Maps</span>
              </a>
            </li>
            <li>
              <a href="tabs.html">
                <i class="bi bi-slash-square"></i>
                <span class="menu-text">Tabs</span>
              </a>
            </li>
            <li>
              <a href="modals.html">
                <i class="bi bi-terminal"></i>
                <span class="menu-text">Modals</span>
              </a>
            </li>
            <li>
              <a href="icons.html">
                <i class="bi bi-textarea"></i>
                <span class="menu-text">Icons</span>
              </a>
            </li>
            <li>
              <a href="typography.html">
                <i class="bi bi-explicit"></i>
                <span class="menu-text">Typography</span>
              </a>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-upc-scan"></i>
                <span class="menu-text">Login/Signup</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="login.html">Login</a>
                </li>
                <li>
                  <a href="signup.html">Signup</a>
                </li>
                <li>
                  <a href="forgot-password.html">Forgot Password</a>
                </li>
                <li>
                  <a href="reset-password.html">Reset Password</a>
                </li>
                <li>
                  <a href="lock-screen.html">Lock Screen</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="page-not-found.html">
                <i class="bi bi-exclamation-diamond"></i>
                <span class="menu-text">Page Not Found</span>
              </a>
            </li>
            <li>
              <a href="maintenance.html">
                <i class="bi bi-exclamation-octagon"></i>
                <span class="menu-text">Maintenance</span>
              </a>
            </li>
            <li class="treeview">
              <a href="#!">
                <i class="bi bi-code-square"></i>
                <span class="menu-text">Nested Menu</span>
              </a>
              <ul class="treeview-menu">
                <li>
                  <a href="#!">Nested 1</a>
                </li>
                <li>
                  <a href="#!">
                    Nested 2
                    <i class="bi bi-caret-right-fill"></i>
                  </a>
                  <ul class="treeview-menu">
                    <li>
                      <a href="#!">Nested 2.1</a>
                    </li>
                    <li>
                      <a href="#!"
                        >Nested 2.2
                        <i class="bi bi-caret-right-fill"></i>
                      </a>
                      <ul class="treeview-menu">
                        <li>
                          <a href="#!">Nested 2.2.1</a>
                        </li>
                        <li>
                          <a href="#!">Nested 2.2.2</a>
                        </li>
                      </ul>
                    </li>
                  </ul>
                </li>
              </ul>
            </li>
          </ul>
        </div>
        <!-- Sidebar menu ends -->
      </nav>
      <!-- Sidebar wrapper end -->

      <!-- App container starts -->
      <div class="app-container">
        <!-- App header starts -->
        <div class="app-header">
          <!-- Toggle buttons start -->
          <div class="d-flex">
            <button class="btn btn-outline-info btn-sm me-3 toggle-sidebar" id="toggle-sidebar">
              <i class="bi bi-list fs-5"></i>
            </button>
            <button class="btn btn-outline-info btn-sm me-3 pin-sidebar" id="pin-sidebar">
              <i class="bi bi-list fs-5"></i>
            </button>
          </div>
          <!-- Toggle buttons end -->

          <!-- App brand start -->
          <div class="app-brand-sm">
            <a href="index.html" class="d-lg-none d-md-block">
              <img src="/assets/images/logo-sm.svg" class="logo" alt="Bootstrap Gallery" />
            </a>
          </div>
          <!-- App brand end -->

          <!-- App header actions start -->
          <div class="header-actions gap-3">
            <!-- Search container starts -->
            <div class="search-container d-lg-flex d-none">
              <input type="text" class="form-control" id="searchData" placeholder="Search" />
              <i class="bi bi-search"></i>
            </div>
            <!-- Search container ends -->

            <!-- Header action starts -->
            <div class="header-actions-block rounded-5 p-2 gap-2 d-sm-flex d-none">
              <div class="dropdown">
                <a
                  class="dropdown-toggle action-icon"
                  href="#!"
                  role="button"
                  data-bs-toggle="dropdown"
                  aria-expanded="false"
                >
                  <i class="bi bi-clipboard-check"></i>
                </a>
                <div class="dropdown-menu dropdown-menu-end dropdown-menu-md">
                  <h5 class="p-3 pb-0">Activity</h5>
                  <!-- Activity start -->
                  <div class="scroll300">
                    <div class="header-activity d-grid p-3">
                      <div class="activity-details info">
                        <div class="ms-4">
                          <p class="mb-1">2 hours ago</p>
                          <p>Messages accepted with attachments.</p>
                        </div>
                      </div>
                      <div class="activity-details danger">
                        <div class="ms-4">
                          <p class="mb-1">3 hours ago</p>
                          <p>Send email notifications of subscriptions and deletions to list owner.</p>
                        </div>
                      </div>
                      <div class="activity-details success">
                        <div class="ms-4">
                          <p class="mb-1">6 hours ago</p>
                          <p>Required change logs activity reports.</p>
                        </div>
                      </div>
                      <div class="activity-details warning">
                        <div class="ms-4">
                          <p class="mb-1">9 hours ago</p>
                          <p>Strategic partnership plan.</p>
                        </div>
                      </div>
                      <div class="activity-details info">
                        <div class="ms-4">
                          <p class="mb-1">3 days ago</p>
                          <p>Send email notifications of subscriptions and deletions to list product manager.</p>
                        </div>
                      </div>
                      <div class="activity-details danger">
                        <div class="ms-4">
                          <p class="mb-1">5 days ago</p>
                          <p>Sharing change logs and activity reports with customer.</p>
                        </div>
                      </div>
                    </div>
                  </div>
                  <!-- Activity end -->
                </div>
              </div>
              <div class="dropdown">
                <a
                  class="dropdown-toggle action-icon"
                  href="#!"
                  role="button"
                  data-bs-toggle="dropdown"
                  aria-expanded="false"
                >
                  <i class="bi bi-handbag"></i>
                  <span class="count-label info">6</span>
                </a>
                <div class="dropdown-menu dropdown-menu-end dropdown-menu-md">
                  <div class="d-grid gap-2 p-3">
                    <div class="p-3 border rounded-2">
                      <h3 class="mb-1">$35,000</h3>
                      <div class="mb-1 d-flex justify-content-between">
                        <span>Revenue</span>
                        <span class="text-primary">+2%</span>
                      </div>
                      <div class="progress small">
                        <div
                          class="progress-bar bg-primary"
                          role="progressbar"
                          style="width: 25%"
                          aria-valuenow="25"
                          aria-valuemin="0"
                          aria-valuemax="100"
                        ></div>
                      </div>
                    </div>
                    <div class="p-3 border rounded-2">
                      <h3 class="mb-1">$48,000</h3>
                      <div class="mb-1 d-flex justify-content-between">
                        <span>Income</span>
                        <span class="text-primary">+7%</span>
                      </div>
                      <div class="progress small">
                        <div
                          class="progress-bar bg-primary"
                          role="progressbar"
                          style="width: 50%"
                          aria-valuenow="50"
                          aria-valuemin="0"
                          aria-valuemax="100"
                        ></div>
                      </div>
                    </div>
                    <div class="p-3 border rounded-2">
                      <h3 class="mb-1">2800</h3>
                      <div class="mb-1 d-flex justify-content-between">
                        <span>Sales</span>
                        <span class="text-danger">+3%</span>
                      </div>
                      <div class="progress small">
                        <div
                          class="progress-bar bg-danger"
                          role="progressbar"
                          style="width: 75%"
                          aria-valuenow="75"
                          aria-valuemin="0"
                          aria-valuemax="100"
                        ></div>
                      </div>
                    </div>
                    <a href="javascript:void(0)" class="btn btn-primary">View all</a>
                  </div>
                </div>
              </div>
              <div class="dropdown">
                <a
                  class="dropdown-toggle action-icon"
                  href="#!"
                  role="button"
                  data-bs-toggle="dropdown"
                  aria-expanded="false"
                >
                  <i class="bi bi-exclamation-triangle"></i>
                  <span class="count-label danger">9</span>
                </a>
                <div class="dropdown-menu dropdown-menu-end dropdown-menu-md">
                  <div class="d-grid gap-4 p-3">
                    <div class="d-flex">
                      <img src="/assets/images/user.png" class="img-3x me-3 rounded-2" alt="Dark Admin Template" />
                      <div class="m-0">
                        <h6>Alexandra Santos</h6>
                        <p class="mb-2">Membership has been ended.</p>
                        <span class="badge bg-danger">Important</span>
                      </div>
                    </div>
                    <div class="d-flex">
                      <img src="/assets/images/user2.png" class="img-3x me-3 rounded-2" alt="Dark Admin Template" />
                      <div class="m-0">
                        <h6>Oliver Hardin</h6>
                        <p class="mb-1">Congratulate, James for new job.</p>
                        <span class="badge bg-info">Remainder</span>
                      </div>
                    </div>
                    <div class="d-flex">
                      <img src="/assets/images/user1.png" class="img-3x me-3 rounded-2" alt="Dark Admin Template" />
                      <div class="m-0">
                        <h6>Murray Patrick</h6>
                        <p class="mb-1">Lewis added new schedule release.</p>
                        <span class="badge bg-success">Network</span>
                      </div>
                    </div>
                    <a href="javascript:void(0)" class="btn btn-primary">View all</a>
                  </div>
                </div>
              </div>
            </div>
            <!-- Header action ends -->

            <!-- Header settings starts -->
            <div class="dropdown">
              <a
                id="userSettings"
                class="dropdown-toggle d-flex py-2 align-items-center text-decoration-none"
                href="#!"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
                <div class="icon-box md bg-info text-white rounded-5">YM</div>
              </a>
              <div class="dropdown-menu dropdown-menu-end">
                <a class="dropdown-item d-flex align-items-center" href="profile.html"
                  ><i class="bi bi-person fs-4 me-2"></i>Profile</a
                >
                <a class="dropdown-item d-flex align-items-center" href="settings.html"
                  ><i class="bi bi-gear fs-4 me-2"></i>Settings</a
                >
                <a class="dropdown-item d-flex align-items-center" href="login.html"
                  ><i class="bi bi-escape fs-4 me-2"></i>Logout</a
                >
              </div>
            </div>
            <!-- Header settings starts -->
          </div>
          <!-- App header actions end -->
        </div>
        <!-- App header ends -->

        <!-- App hero header starts -->
        <div class="app-hero-header d-flex align-items-center">
          <!-- Breadcrumb start -->
          <ol class="breadcrumb d-none d-lg-flex">
            <li class="breadcrumb-item">
              <i class="bi bi-house lh-1"></i>
              <a href="index.html" class="text-decoration-none">Home</a>
            </li>
            <li class="breadcrumb-item" aria-current="page">Dashboard</li>
          </ol>
          <!-- Breadcrumb end -->

          <!-- Filter start -->
          <div class="ms-auto d-flex flex-row gap-1 day-filters">
            <button class="btn btn-sm">Today</button>
            <button class="btn btn-sm">7D</button>
            <button class="btn btn-sm">2W</button>
            <button class="btn btn-sm">1M</button>
            <button class="btn btn-sm">3M</button>
            <button class="btn btn-sm">6M</button>
            <button class="btn btn-sm btn-info">1Y</button>
          </div>
          <!-- Filter end -->
        </div>
        <!-- App Hero header ends -->

        <!-- App body starts -->
        <div class="app-body">
          <!-- Row start -->
          <div class="row gx-3">
            <div class="col-xl-3 col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-body">
                  <div class="mb-2">
                    <i class="bi bi-bar-chart fs-1 text-primary lh-1"></i>
                  </div>
                  <div class="d-flex align-items-center justify-content-between">
                    <h5 class="m-0 fw-normal">Sales</h5>
                    <h3 class="m-0 text-primary">3500</h3>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-xl-3 col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-body">
                  <div class="mb-2">
                    <i class="bi bi-bag-check fs-1 text-primary lh-1"></i>
                  </div>
                  <div class="d-flex align-items-center justify-content-between">
                    <h5 class="m-0 fw-normal">Orders</h5>
                    <h3 class="m-0 text-primary">2900</h3>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-xl-3 col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-body">
                  <div class="arrow-label">+18%</div>
                  <div class="mb-2">
                    <i class="bi bi-box-seam fs-1 text-primary lh-1"></i>
                  </div>
                  <div class="d-flex align-items-center justify-content-between">
                    <h5 class="m-0 fw-normal">Items</h5>
                    <h3 class="m-0 text-primary">6500</h3>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-xl-3 col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-body">
                  <div class="arrow-label">+21%</div>
                  <div class="mb-2">
                    <i class="bi bi-check-circle fs-1 text-danger lh-1"></i>
                  </div>
                  <div class="d-flex align-items-center justify-content-between">
                    <h5 class="m-0 fw-normal">Signups</h5>
                    <h3 class="m-0 text-danger">7200</h3>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Row end -->

          <!-- Row start -->
          <div class="row gx-3">
            <div class="col-xxl-12">
              <div class="card mb-3">
                <div class="card-header d-flex justify-content-between align-items-center">
                  <h5 class="card-title">Overview</h5>
                  <button class="btn btn-outline-primary btn-sm ms-auto">Download</button>
                </div>
                <div class="card-body">
                  <!-- Row start -->
                  <div class="row gx-3">
                    <div class="col-lg-5 col-sm-12 col-12">
                      <h6 class="text-center mb-3">Visitors</h6>
                      <div id="visitors"></div>
                      <div class="my-3 text-center">
                        <div class="badge bg-danger bg-opacity-10 text-danger">10% higher than last month</div>
                      </div>
                    </div>
                    <div class="col-lg-2 col-sm-12 col-12">
                      <div class="border rounded-4 px-2 py-4 h-100 text-center">
                        <h6 class="mt-3 mb-5">Monthly Average</h6>
                        <div class="mb-5">
                          <h2 class="text-primary">9600</h2>
                          <h6>Visitors</h6>
                        </div>
                        <div class="mb-4">
                          <h2 class="text-danger">$450<sup>k</sup></h2>
                          <h6>Sales</h6>
                        </div>
                      </div>
                    </div>
                    <div class="col-lg-5 col-sm-12 col-12">
                      <h6 class="text-center mb-3">Sales</h6>
                      <div id="sales"></div>
                      <div class="my-3 text-center">
                        <div class="badge bg-primary bg-opacity-10 text-primary">12% higher than last month</div>
                      </div>
                    </div>
                  </div>
                  <!-- Row ends -->
                </div>
              </div>
            </div>
          </div>
          <!-- Row ends -->

          <!-- Row start -->
          <div class="row gx-3">
            <div class="col-xl-8 col-lg-12">
              <div class="card mb-3">
                <div class="card-header">
                  <h5 class="card-title">Team Activity</h5>
                </div>
                <div class="card-body">
                  <ul class="m-0 p-0">
                    <li class="team-activity d-flex flex-wrap">
                      <div class="activity-time py-2 me-3">
                        <p class="m-0">10:30AM</p>
                        <span class="badge bg-primary">New</span>
                      </div>
                      <div class="d-flex flex-column py-2">
                        <h6>Earth - Admin Dashboard</h6>
                        <p class="m-0">by Elnathan Lois</p>
                      </div>
                      <div class="ms-auto mt-4">
                        <div class="progress small mb-1">
                          <div
                            class="progress-bar"
                            role="progressbar"
                            style="width: 25%"
                            aria-valuenow="25"
                            aria-valuemin="0"
                            aria-valuemax="100"
                          ></div>
                        </div>
                        <p>(225 of 700gb)</p>
                      </div>
                    </li>
                    <li class="team-activity d-flex flex-wrap">
                      <div class="activity-time py-2 me-3">
                        <p class="m-0">11:30AM</p>
                        <span class="badge bg-primary">Task</span>
                      </div>
                      <div class="d-flex flex-column py-2">
                        <h6>Bootstrap Gallery Admin Templates</h6>
                        <p class="m-0">by Patrobus Nicole</p>
                      </div>
                      <div class="ms-auto mt-4">
                        <div class="progress small mb-1">
                          <div
                            class="progress-bar"
                            role="progressbar"
                            style="width: 80%"
                            aria-valuenow="80"
                            aria-valuemin="0"
                            aria-valuemax="100"
                          ></div>
                        </div>
                        <p>90% completed</p>
                      </div>
                    </li>
                    <li class="team-activity d-flex flex-wrap">
                      <div class="activity-time py-2 me-3">
                        <p class="m-0">12:50PM</p>
                        <span class="badge bg-danger">Closed</span>
                      </div>
                      <div class="d-flex flex-column py-2">
                        <h6>Bootstrap Admin Themes</h6>
                        <p class="m-0">by Abilene Omega</p>
                      </div>
                      <div class="ms-auto mt-3">
                        <div id="sparkline1"></div>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div class="col-xl-4 col-lg-12">
              <div class="card mb-3">
                <div class="card-header">
                  <h5 class="card-title">Tasks</h5>
                </div>
                <div class="card-body">
                  <div class="auto-align-graph">
                    <div id="tasks"></div>
                  </div>
                  <div class="grid text-center">
                    <div class="g-col-4">
                      <i class="bi bi-triangle text-danger"></i>
                      <h3 class="m-0 mt-1">6</h3>
                      <p class="m-0">New</p>
                    </div>
                    <div class="g-col-4">
                      <i class="bi bi-triangle text-primary"></i>
                      <h3 class="m-0 mt-1 fw-bolder">9</h3>
                      <p class="m-0">Pending</p>
                    </div>
                    <div class="g-col-4">
                      <i class="bi bi-triangle text-success"></i>
                      <h3 class="m-0 mt-1">12</h3>
                      <p class="m-0">Completed</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Row end -->

          <!-- Row start -->
          <div class="row gx-3">
            <div class="col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-header">
                  <h5 class="card-title">Stats</h5>
                </div>
                <div class="card-body">
                  <div class="scroll350">
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill text-primary me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-bag fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2">You have spent about <b>65%</b> of your annual budget.</p>
                      <div class="ms-auto badge bg-primary bg-opacity-10 text-primary small">24/12/2023</div>
                    </div>
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill text-primary me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-check-circle fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2">New admin dashboard purchased, and payment paid through online.</p>
                      <div class="ms-auto badge bg-primary bg-opacity-10 text-primary small">23/12/2023</div>
                    </div>
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill text-primary me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-clipboard-check fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2">A new ticket opened and assigned to <b>Zion</b>.</p>
                      <div class="ms-auto badge bg-primary bg-opacity-10 text-primary small">22/12/2023</div>
                    </div>
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-slash-circle fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2">Thanks <b>Sarah</b>, I want you to share Jim's profile.</p>
                      <div class="ms-auto badge bg-primary bg-opacity-10 text-primary small">21/12/2023</div>
                    </div>
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-envelope-open fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2"><b>Ora Mahoney,</b> has completed the design of the CRM admin application.</p>
                      <div class="ms-auto badge bg-danger bg-opacity-10 text-danger small">20/12/2023</div>
                    </div>
                    <div class="d-flex align-items-center mb-4">
                      <i class="bi bi-record-fill me-2"></i>
                      <div class="d-flex p-2 bg-primary rounded-circle me-3">
                        <i class="bi bi-envelope-open fs-4 text-white lh-1"></i>
                      </div>
                      <p class="m-0 me-2"><b>Daren Boyd,</b> received the order.</p>
                      <div class="ms-auto badge bg-danger bg-opacity-10 text-danger small">18/12/2023</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-sm-6 col-12">
              <div class="card mb-3">
                <div class="card-header">
                  <h5 class="card-title">Orders</h5>
                </div>
                <div class="card-body">
                  <div class="row gx-3">
                    <div class="col-lg-6 col-sm-6 col-12">
                      <div class="d-flex flex-column align-items-start">
                        <div class="icon-box lg border border-primary rounded-5 my-3">
                          <span class="bi bi-bar-chart-line fs-2 text-primary lh-1"></span>
                        </div>
                        <h3 class="mt-4">$780<sup>k</sup></h3>
                        <p class="mb-4">Highest sales growth in last two years.</p>
                        <div class="d-flex p-3 flex-column border border-primary rounded-2 mb-4">
                          <h6>Target - <span>75/100</span></h6>
                          <div class="progress small">
                            <div
                              class="progress-bar bg-primary"
                              role="progressbar"
                              style="width: 75%"
                              aria-valuenow="75"
                              aria-valuemin="0"
                              aria-valuemax="100"
                            ></div>
                          </div>
                        </div>
                        <small class="badge bg-primary bg-opacity-10 text-primary rounded-1">47% High growth</small>
                      </div>
                    </div>
                    <div class="col-lg-6 col-sm-6 col-12">
                      <div id="orders"></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-12">
              <div class="card mb-3">
                <div class="card-header">
                  <h5 class="card-title">Sales in USA</h5>
                </div>
                <div class="card-body">
                  <!-- Chart starts -->
                  <div id="us-map4" class="chart-height-xxl"></div>
                  <!-- Chart ends -->

                  <div class="grid gap-3 mt-4">
                    <div class="g-col-4">
                      <div class="p-3 flex-column border border-primary rounded-2">
                        <p class="mb-1">Total Sales</p>
                        <h4 class="fw-bold mb-2">$9,800</h4>
                        <div class="progress small">
                          <div
                            class="progress-bar bg-primary"
                            role="progressbar"
                            style="width: 75%"
                            aria-valuenow="75"
                            aria-valuemin="0"
                            aria-valuemax="100"
                          ></div>
                        </div>
                      </div>
                    </div>
                    <div class="g-col-4">
                      <div class="p-3 flex-column border border-primary rounded-2">
                        <p class="mb-1">Active Users</p>
                        <h4 class="fw-bold mb-2">8,900</h4>
                        <div class="progress small">
                          <div
                            class="progress-bar bg-success"
                            role="progressbar"
                            style="width: 75%"
                            aria-valuenow="75"
                            aria-valuemin="0"
                            aria-valuemax="100"
                          ></div>
                        </div>
                      </div>
                    </div>
                    <div class="g-col-4">
                      <div class="p-3 flex-column border border-primary rounded-2">
                        <p class="mb-1">Daily Users</p>
                        <h4 class="fw-bold mb-2">3,600</h4>
                        <div class="progress small">
                          <div
                            class="progress-bar bg-danger"
                            role="progressbar"
                            style="width: 75%"
                            aria-valuenow="75"
                            aria-valuemin="0"
                            aria-valuemax="100"
                          ></div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Row ends -->
        </div>
        <!-- App body ends -->

        <!-- App footer start -->
        <div class="app-footer">
          <span>© Bootstrap Gallery 2024</span>
        </div>
        <!-- App footer end -->
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <!-- <div id="thirdPartyScripts"></div> -->
</template>
