
-- +migrate Up

DROP TRIGGER trg_ensure_alert_dedup_key ON public.alerts;
DROP FUNCTION fn_ensure_alert_dedup_key();

-- +migrate Down

-- +migrate StatementBegin
CREATE FUNCTION fn_ensure_alert_dedup_key() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.dedup_key ISNULL THEN
        NEW.dedup_key = 
            concat(
                'auto:1:',
                encode(digest(concat(NEW."description"), 'sha512'), 'hex')
            );
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';
-- +migrate StatementEnd

CREATE TRIGGER trg_ensure_alert_dedup_key
BEFORE INSERT ON public.alerts
FOR EACH ROW WHEN ((new.status <> 'closed'::enum_alert_status))
EXECUTE PROCEDURE fn_ensure_alert_dedup_key();
