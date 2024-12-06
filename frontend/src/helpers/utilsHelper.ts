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

export function calculateTotalHelper(characteristics: any[]): number {
  return characteristics.reduce((total, characteristic) => {
    const value = parseInt(characteristic?.score) || 0;
    return total + value;
  }, 0);
}

export function allRadiosFilledHelper(characteristic: any[]): boolean {
  return characteristic.every(char => char.score !== undefined && char.score !== null && char.score !== 0);
}

export function updateTotalScoreHelper(document: any, totalScoreRef: Ref<number, number>) {
  const rows = document.querySelectorAll('tbody tr');
  let total = 0;
  rows.forEach((row: HTMLTableRowElement) => {
    const lastTd = row.querySelector('td:last-child span');
    if (lastTd) {
      total += parseFloat(lastTd.textContent || '0');
    }
  });
  totalScoreRef.value = parseFloat(total.toFixed(2));
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
