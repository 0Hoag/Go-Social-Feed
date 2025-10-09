# Go-Social-Feed

SUCCESS - Basic Flow
TC01: Happy path với mix events (non-recurring + recurring)
TC02: Empty result (no events in date range)
TC03: Specific event IDs filter
TC04: User permissions (creator/assigned/public/system/branch/department access)
TC05: Date range boundaries và time zone handling
TC06: All recurring patterns (daily/weekly/monthly/yearly) với RepeatUntil
TC07: Filter instances theo branch/department/categories/assign logic
TC08: Sort by start time functionality
SUCCESS - Edge Cases
TC09: Monthly events với ngày không tồn tại (31/2, 30/2, etc.)
TC10: Yearly events với leap year (29/2)
TC11: Existing recurring instances + newly generated instances
TC12: Multiple timezones, categories, users trong result
TC13: User có BranchPlusIds và DepartmentPlusIds
TC14: Events với excluded categories nhưng user vẫn access được
TC15: Mix all-day và timed events
ERROR - Dependencies
TC16: shopUC.GetSessionUser fails
TC17: elementUC.ListShopElement fails (excluded categories)
TC18: repo.List fails (non-recurring events)
TC19: getRecurringInstanceInDateRange fails
TC20: Concurrent metadata loading fails (eventcategoryUC.List, elementUC.List, shopUC.ListAllUsers)
TC21: errgroup.Wait() returns error
ERROR - GetRecurringInstanceInDateRange
TC22: GetGenRTsInDateRange hoặc GetGenRTsNotInDateRange fails
TC23: ListByIDs fails (events not found)
TC24: elementUC.List fails (timezones not found)
TC25: generateInstancesForMonth fails
TC26: CreateRecurringTracking fails
TC27: CreateManyRecurringInstances fails
TC28: getExistingInstances fails
TC29: ListRecurringInstancesByEventIDs fails
ERROR - Business Logic
TC30: Invalid date range (start > end)
TC31: Infinite loop trong getNextOccurrence
TC32: Hit MaxInstancesPerMonth limit
TC33: Event start time after generation range
TC34: Data inconsistency (event/timezone/category/user not found)
TC35: Filter logic với invalid conditions
PERFORMANCE & BOUNDARY
TC36: Large dataset (1000+ events, 10000+ instances)
TC37: Date range span multiple years
TC38: Memory/resource constraints
TC39: Concurrent access và race conditions
TC40: Database connection issues
INPUT VALIDATION
TC41: Invalid/null context
TC42: Invalid scope
TC43: Invalid event IDs format
TC44: Empty/null input parameters
