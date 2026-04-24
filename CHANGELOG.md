## 1.6.0 (Unreleased)

NEW FEATURES:

* `terraform test`: The `terraform test` command now supports a `-junit-xml=FILE` option that, if specified, will write a JUnit XML test results file to the given location after running the tests. ([#34291](https://github.com/opentofu/opentofu/issues/34291))

ENHANCEMENTS:

* `terraform init`: Improved error messages when a required provider is not available in the configured provider registry. ([#34102](https://github.com/opentofu/opentofu/issues/34102))
* config: The `templatefile` function now supports recursive calls, allowing templates to include other templates. ([#34211](https://github.com/opentofu/opentofu/issues/34211))
* backend/s3: Added support for the `use_path_style` argument to force path-style S3 API calls, which is required for some S3-compatible services. ([#34088](https://github.com/opentofu/opentofu/issues/34088))

BUG FIXES:

* `terraform plan`: Fixed a panic that could occur when planning a resource with a `dynamic` block inside a `lifecycle` `precondition` or `postcondition`. ([#34309](https://github.com/opentofu/opentofu/issues/34309))
* `terraform apply`: Fixed an issue where `create_before_destroy` replacements could fail when the prior state contained unknown values from a previous plan. ([#34255](https://github.com/opentofu/opentofu/issues/34255))
* config: Fixed incorrect type constraint inference for `object` type expressions that include optional attributes with default values. ([#34190](https://github.com/opentofu/opentofu/issues/34190))

## 1.5.7 (September 27, 2023)

BUG FIXES:

* `terraform init`: Fixed a crash when a module call uses an invalid version constraint string. ([#34111](https://github.com/opentofu/opentofu/issues/34111))
* backend/gcs: Fixed an issue where the GCS backend would not correctly handle OAuth2 token refresh. ([#34098](https://github.com/opentofu/opentofu/issues/34098))

## 1.5.6 (September 14, 2023)

BUG FIXES:

* `terraform test`: Fixed an issue where test files using `run` blocks with `command = plan` could produce incorrect results when the plan contained resource deletions. ([#34008](https://github.com/opentofu/opentofu/issues/34008))
* config: Fixed a crash when using `for_each` with a `null` value in certain contexts. ([#34021](https://github.com/opentofu/opentofu/issues/34021))
* `terraform apply`: Corrected an edge case where `depends_on` between modules could cause unnecessary re-evaluation of data sources. ([#33987](https://github.com/opentofu/opentofu/issues/33987))

## Previous Releases

<!-- Note to self: when cutting a new release, remember to update the version links below to point to the correct tags. -->

For information on prior major and minor releases, see their changelogs:

* [v1.4](https://github.com/opentofu/opentofu/blob/v1.4/CHANGELOG.md)
* [v1.3](https://github.com/opentofu/opentofu/blob/v1.3/CHANGELOG.md)
* [v1.2](https://github.com/opentofu/opentofu/blob/v1.2/CHANGELOG.md)
* [v1.1](https://github.com/opentofu/opentofu/blob/v1.1/CHANGELOG.md)
* [v1.0](https://github.com/opentofu/opentofu/blob/v1.0/CHANGELOG.md)
* [v0.15](https://github.com/opentof