import {
  Link,
  PublicFieldProps,
  InjectedFieldProps,
  sanitizeFieldRestProps,
} from "react-admin";

import {
  useReferenceManyFieldController,
  useRecordContext,
  useTimeout,
  useCreatePath,
} from "ra-core";

import { Typography, TypographyProps, CircularProgress } from "@mui/material";
import ErrorIcon from "@mui/icons-material/Error";

export const CustomManyCount = (props: CustomManyCountProps) => {
  const {
    reference,
    target,
    filter,
    link,
    resource,
    sortBy = "id",
    sortByOrder = "ASC",
    source = "id",
    timeout = 1000,
    ...rest
  } = props;

  const record = useRecordContext(props);
  const oneSecondHasPassed = useTimeout(timeout);
  const createPath = useCreatePath();

  const { isLoading, error, total } = useReferenceManyFieldController({
    filter,
    page: 1,
    perPage: 1,
    record,
    reference,
    resource,
    source,
    target,
    sort: { field: sortBy, order: sortByOrder },
  });

  const body = isLoading ? (
    oneSecondHasPassed ? (
      <CircularProgress size={14} />
    ) : (
      ""
    )
  ) : error ? (
    <ErrorIcon color="error" fontSize="small" titleAccess="error" />
  ) : (
    total
  );

  return link ? (
    // @ts-ignore TypeScript complains that the props for <a> aren't the same as for <span>
    <Link
      to={{
        pathname: createPath({ resource: reference, type: "list" }),
        search: `filter=${JSON.stringify({
          ...(filter || {}),
          [target]: record[source],
        })}`,
      }}
      variant="body2"
      component="span"
      onClick={(e) => e.stopPropagation()}
      {...rest}
    >
      {body}
    </Link>
  ) : (
    <Typography
      component="span"
      variant="body2"
      {...sanitizeFieldRestProps(rest)}
    >
      {body}
    </Typography>
  );
};

export interface CustomManyCountProps
  extends PublicFieldProps,
    InjectedFieldProps,
    Omit<TypographyProps, "textAlign"> {
  reference: string;
  target: string;
  filter?: any;
  label?: string;
  link?: boolean;
  resource?: string;
  source?: string;
  timeout?: number;
}
