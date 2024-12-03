import { Reactive, Ref } from 'vue';
import postscribe from 'postscribe';

export function formValidationAreAllFieldsFilled(reactiveFormObject: Reactive<any>, formClassRef: Ref): boolean {
  for (const key in reactiveFormObject) {
    const value = reactiveFormObject[key as keyof typeof reactiveFormObject];
    if (!value) {
      formClassRef.value = 'was-validated';
      return false;
    }
  }
  return true;
}

export async function loadThirdPartyJS(jsFilePaths: string[]) {
  for (const filepath of jsFilePaths) {
    await postscribe('#thirdPartyScripts', `<script src="${filepath}"><\/script>`);
  }
}

export function generateCalendarDateString() {
  return new Date().toISOString().split('T')[0];
}
