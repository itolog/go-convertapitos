import type { AxiosError } from "axios";
import type {
  ApiResponseError,
  ValidationErrorFields,
} from "@/generated/apiClient/data-contracts.ts";
import type { Vueform } from "@vueform/vueform";

export const handleFormError = (
  data: AxiosError<ApiResponseError>,
  details: { type: string },
  form$: Vueform,
) => {
  const error = data.response?.data.error;
  // @ts-expect-error - Vueform docs show clear() can be called without arguments
  form$.messageBag.clear(); // clear message bag

  switch (details.type) {
    // Error occured while preparing elements (no submit happened)
    case "prepare":
      // @ts-expect-error - Vueform docs show append() can be called without arguments
      form$.messageBag.append("Could not prepare form");
      break;

    // Error occured because response status is outside of 2xx
    case "submit":
      if (error?.fields?.length) {
        error.fields.forEach((item: ValidationErrorFields) => {
          const field = item?.field?.toLowerCase() || "";
          const messageBag = form$.el$(field)?.messageBag;
          // @ts-expect-error - Vueform docs show clear() can be called without arguments
          messageBag?.clear();
          // @ts-expect-error - Vueform docs show append() can be called without arguments
          messageBag?.append(`${field}: ${item.tag} ${item.param}`);
        });
      }
      // @ts-expect-error - Vueform docs show append() can be called without arguments
      form$.messageBag.append(error?.message);
      break;

    // Request cancelled (no response object)
    case "cancel":
      // @ts-expect-error - Vueform docs show append() can be called without arguments
      form$.messageBag.append("Request cancelled");
      break;

    // Some other errors happened (no response object)
    case "other":
      // @ts-expect-error - Vueform docs show append() can be called without arguments
      form$.messageBag.append("Couldn't submit form");
      break;
  }
};

export const clearFormErrors = (form$: Vueform) => {
  // @ts-expect-error - Vueform docs show clear() can be called without arguments
  form$.messageBag.clear();

  Object.keys(form$.elements$).forEach((fieldName) => {
    if (fieldName !== "button" && fieldName !== "submit") {
      const field = form$.el$(fieldName);
      if (field && field.messageBag) {
        // @ts-expect-error - Vueform docs show clear() can be called without arguments
        field.messageBag.clear();
      }
    }
  });
};
