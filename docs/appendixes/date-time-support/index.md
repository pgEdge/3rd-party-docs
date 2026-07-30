<a id="datetime-appendix"></a>

# Date/Time Support

 PostgreSQL uses an internal heuristic parser for all date/time input support. Dates and times are input as strings, and are broken up into distinct fields with a preliminary determination of what kind of information can be in the field. Each field is interpreted and either assigned a numeric value, ignored, or rejected. The parser contains internal lookup tables for all textual fields, including months, days of the week, and time zones.

 This appendix includes information on the content of these lookup tables and describes the steps used by the parser to decode dates and times.

- [Date/Time Input Interpretation](date-time-input-interpretation.md#datetime-input-rules)
- [Handling of Invalid or Ambiguous Timestamps](handling-of-invalid-or-ambiguous-timestamps.md#datetime-invalid-input)
- [Date/Time Key Words](date-time-key-words.md#datetime-keywords)
- [Date/Time Configuration Files](date-time-configuration-files.md#datetime-config-files)
- [POSIX Time Zone Specifications](posix-time-zone-specifications.md#datetime-posix-timezone-specs)
- [History of Units](history-of-units.md#datetime-units-history)
- [Julian Dates](julian-dates.md#datetime-julian-dates)
