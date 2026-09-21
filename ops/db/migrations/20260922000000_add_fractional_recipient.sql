-- Modify "rules" table
ALTER TABLE "rules" ADD COLUMN "fractional_recipient" integer NULL, ADD CONSTRAINT "rules_fractional_recipient_check" CHECK (fractional_recipient = ANY (ARRAY[1, 2])), ADD CONSTRAINT "chk_fractional_recipient" CHECK (((fractional_calculation = 1) AND (fractional_recipient IS NULL)) OR ((fractional_calculation <> 1) AND (fractional_recipient IS NOT NULL)));
-- Set comment to column: "fractional_recipient" on table: "rules"
COMMENT ON COLUMN "rules"."fractional_recipient" IS '1: 1位の人, 2: 最下位の人';
-- Modify "game_rules" table
ALTER TABLE "game_rules" ADD COLUMN "fractional_recipient" integer NULL, ADD CONSTRAINT "game_rules_fractional_recipient_check" CHECK (fractional_recipient = ANY (ARRAY[1, 2]));
-- Set comment to column: "fractional_recipient" on table: "game_rules"
COMMENT ON COLUMN "game_rules"."fractional_recipient" IS '1: 1位の人, 2: 最下位の人';
