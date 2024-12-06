import { Reactive, Ref } from 'vue';
import postscribe from 'postscribe';
import { RouteLocationNormalizedLoadedGeneric } from 'vue-router';
export function currentYearHelper() {
  return new Date().getFullYear();
}

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

export async function loadThirdPartyJSHelper(jsFilePaths: string[]) {
  for (const filepath of jsFilePaths) {
    await postscribe('#thirdPartyScripts', `<script src="${filepath}"><\/script>`);
  }
}

export function generateCalendarDateStringHelper() {
  return new Date().toISOString().split('T')[0];
}

export function hideNavSidebarHelper(
  route: RouteLocationNormalizedLoadedGeneric,
  hideNavSidebarRef: Ref<boolean, boolean>,
) {
  // TODO: Maybe make this a use an array
  if (route.path.includes('appraisal')) {
    hideNavSidebarRef.value = true;
  } else {
    hideNavSidebarRef.value = false;
  }
}

export function breadCrumbPageLinkPreviousHelper(
  route: RouteLocationNormalizedLoadedGeneric,
  authStore: any,
): Record<string, string> {
  const user = authStore.getState;
  const returnProfile = { link: `/user/${user.userId}`, linkName: 'Profile' };
  const returnDashboard = { link: '/dashboard', linkName: 'Dashboard' };
  if (route.path === '/dashboard') return { link: '#', linkName: ' ' };
  if (user.userPrivilegeLevel === 'e') return returnDashboard;
  return returnProfile;
}
