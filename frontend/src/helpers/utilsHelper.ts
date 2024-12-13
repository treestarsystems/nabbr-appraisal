import { Reactive, Ref } from 'vue';
import postscribe from 'postscribe';
import { RouteLocationNormalizedLoadedGeneric } from 'vue-router';
export function currentYearHelper() {
  return new Date().getFullYear();
}

/**
 * Check if all fields are filled in a form
 * @param {Reactive<any>} reactiveFormObject reactive form object
 * @param {Ref} formClassRef ref form object
 * @returns {boolean} true if all fields are filled, false otherwise
 */

export function formValidationAreAllFieldsFilledHelper(reactiveFormObject: Reactive<any>, formClassRef: Ref): boolean {
  for (const key in reactiveFormObject) {
    const value = reactiveFormObject[key as keyof typeof reactiveFormObject];
    if (!value) {
      formClassRef.value = 'was-validated';
      return false;
    }
  }
  return true;
}

/**
 * Use postscribe to load third party js files
 * @param {string[]} jsFilePaths array of js file paths
 * @returns {Promise<void>}
 */

export async function loadThirdPartyJSHelper(jsFilePaths: string[]): Promise<void> {
  for (const filepath of jsFilePaths) {
    await postscribe('#thirdPartyScripts', `<script src="${filepath}"><\/script>`);
  }
}

/**
 * Generate a calendar date string
 * @returns {string} calendar date string in YYYY-MM-DD format
 */
export function generateCalendarDateStringHelper(): string {
  return new Date().toISOString().split('T')[0];
}

/**
 * Hide the navigation sidebar based on the route path then modify the passed hideNavSidebarRef object
 * @param {RouteLocationNormalizedLoadedGeneric} route vue router route object
 * @param {Ref<boolean, boolean>} hideNavSidebarRef target ref to hide the navigation sidebar
 * @returns {void}
 */

export function hideNavSidebarHelper(
  route: RouteLocationNormalizedLoadedGeneric,
  hideNavSidebarRef: Ref<boolean, boolean>,
) {
  // TODO: Maybe make this a use an array...maybe.
  if (route.path.includes('appraisal')) {
    hideNavSidebarRef.value = true;
  } else {
    hideNavSidebarRef.value = false;
  }
}

/**
 * Create a breadcrumb link for the previous page based on the route path and user privilege level
 * @param {RouteLocationNormalizedLoadedGeneric} route vue router route object
 * @param {Ref<boolean, boolean>} hideNavSidebarRef target ref to hide the navigation sidebar
 * @returns {Record<string, string>} object containing the link and link name
 */

export function breadCrumbPageLinkPreviousHelper(
  route: RouteLocationNormalizedLoadedGeneric,
  authStore: any,
): Record<string, string> {
  const user = authStore.getState;
  const returnProfile = { link: `/user/${user?.userId}`, linkName: 'Profile' };
  const returnDashboard = { link: '/dashboard', linkName: 'Dashboard' };
  if (route.path === '/dashboard') return { link: '#', linkName: ' ' };
  if (user?.userPrivilegeLevel === 'ADMIN') return returnDashboard;
  return returnProfile;
}
